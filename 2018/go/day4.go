package main

import (
	"sort"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const test = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

const input = `[1518-07-10 23:54] Guard #3167 begins shift
[1518-04-15 00:20] falls asleep
[1518-09-30 00:49] wakes up
[1518-05-29 00:57] wakes up
[1518-11-03 00:00] Guard #1319 begins shift
[1518-06-07 00:03] Guard #1619 begins shift
[1518-09-02 00:15] falls asleep
[1518-04-26 00:48] wakes up
[1518-11-02 00:56] wakes up
[1518-08-30 00:22] falls asleep
[1518-06-11 00:00] Guard #1319 begins shift
[1518-08-07 00:57] wakes up
[1518-09-08 00:59] wakes up
[1518-11-23 00:00] Guard #2861 begins shift
[1518-05-21 00:39] wakes up
[1518-08-19 23:57] Guard #1319 begins shift
[1518-07-05 00:39] falls asleep
[1518-06-02 23:57] Guard #101 begins shift
[1518-09-02 00:58] wakes up
[1518-11-08 23:59] Guard #263 begins shift
[1518-10-07 00:15] falls asleep
[1518-06-06 00:03] Guard #1319 begins shift
[1518-08-11 00:57] wakes up
[1518-06-01 00:51] falls asleep
[1518-10-24 00:26] wakes up
[1518-05-16 23:57] Guard #1619 begins shift
[1518-11-17 00:13] falls asleep
[1518-04-21 00:37] wakes up
[1518-09-03 00:55] wakes up
[1518-10-30 00:10] wakes up
[1518-04-18 00:49] wakes up
[1518-11-07 00:51] wakes up
[1518-07-05 00:52] falls asleep
[1518-09-09 00:46] wakes up
[1518-06-06 00:33] falls asleep
[1518-06-08 00:42] wakes up
[1518-08-17 00:02] falls asleep
[1518-11-17 00:24] wakes up
[1518-05-28 00:50] falls asleep
[1518-04-27 00:00] Guard #3203 begins shift
[1518-04-20 00:01] Guard #2539 begins shift
[1518-05-25 00:24] falls asleep
[1518-08-13 00:04] falls asleep
[1518-09-05 23:57] Guard #1913 begins shift
[1518-08-31 00:03] Guard #521 begins shift
[1518-05-05 00:42] falls asleep
[1518-10-26 00:01] Guard #1069 begins shift
[1518-07-20 00:03] Guard #809 begins shift
[1518-07-04 00:24] wakes up
[1518-04-01 00:00] Guard #3167 begins shift
[1518-07-07 00:50] falls asleep
[1518-07-12 00:48] wakes up
[1518-10-02 00:16] falls asleep
[1518-04-10 00:17] falls asleep
[1518-09-17 00:06] falls asleep
[1518-10-23 23:50] Guard #3203 begins shift
[1518-08-20 00:48] falls asleep
[1518-08-09 00:15] wakes up
[1518-11-08 00:19] falls asleep
[1518-10-03 00:56] falls asleep
[1518-11-06 00:21] falls asleep
[1518-04-18 00:03] Guard #421 begins shift
[1518-07-15 00:26] wakes up
[1518-09-29 23:49] Guard #1237 begins shift
[1518-05-28 00:04] Guard #2441 begins shift
[1518-10-10 23:52] Guard #3167 begins shift
[1518-04-29 00:42] wakes up
[1518-08-08 00:01] falls asleep
[1518-06-24 23:58] Guard #2221 begins shift
[1518-04-06 00:09] falls asleep
[1518-09-13 00:01] Guard #101 begins shift
[1518-05-10 00:08] falls asleep
[1518-09-05 00:42] falls asleep
[1518-06-02 00:17] falls asleep
[1518-07-30 00:57] wakes up
[1518-09-27 00:07] wakes up
[1518-06-16 23:59] Guard #89 begins shift
[1518-06-02 00:53] wakes up
[1518-09-24 00:00] Guard #1069 begins shift
[1518-08-16 00:54] wakes up
[1518-06-27 00:46] wakes up
[1518-06-08 00:02] Guard #3167 begins shift
[1518-09-02 00:02] Guard #953 begins shift
[1518-10-29 00:18] falls asleep
[1518-09-15 00:59] wakes up
[1518-04-20 00:23] falls asleep
[1518-07-01 00:20] falls asleep
[1518-05-24 00:43] falls asleep
[1518-09-02 00:50] falls asleep
[1518-06-29 00:17] falls asleep
[1518-06-14 00:04] Guard #881 begins shift
[1518-08-04 00:02] Guard #421 begins shift
[1518-08-16 00:02] Guard #263 begins shift
[1518-07-07 00:17] wakes up
[1518-11-21 00:45] falls asleep
[1518-10-04 00:03] falls asleep
[1518-05-07 00:03] Guard #421 begins shift
[1518-08-30 00:24] wakes up
[1518-11-23 00:29] falls asleep
[1518-04-25 00:00] Guard #2381 begins shift
[1518-09-02 00:45] wakes up
[1518-05-15 00:52] falls asleep
[1518-09-05 00:32] falls asleep
[1518-07-01 23:47] Guard #1069 begins shift
[1518-10-08 00:21] falls asleep
[1518-08-13 00:07] wakes up
[1518-06-16 00:40] wakes up
[1518-05-15 23:56] Guard #1601 begins shift
[1518-07-25 00:00] Guard #2221 begins shift
[1518-04-03 00:53] wakes up
[1518-06-15 00:55] wakes up
[1518-04-27 00:59] wakes up
[1518-08-09 23:53] Guard #421 begins shift
[1518-09-01 00:01] Guard #2381 begins shift
[1518-10-05 00:07] falls asleep
[1518-05-28 00:17] falls asleep
[1518-08-12 23:50] Guard #1601 begins shift
[1518-10-11 00:51] wakes up
[1518-08-16 00:13] falls asleep
[1518-05-21 00:50] wakes up
[1518-04-18 23:56] Guard #3203 begins shift
[1518-06-22 00:34] wakes up
[1518-05-31 00:57] wakes up
[1518-10-19 00:42] wakes up
[1518-06-20 00:53] falls asleep
[1518-09-30 00:26] falls asleep
[1518-05-14 00:56] wakes up
[1518-10-23 00:51] wakes up
[1518-06-17 23:59] Guard #421 begins shift
[1518-05-05 00:37] wakes up
[1518-06-12 00:07] falls asleep
[1518-07-09 00:41] falls asleep
[1518-05-13 00:39] falls asleep
[1518-09-04 00:38] wakes up
[1518-09-10 00:00] Guard #881 begins shift
[1518-10-11 00:01] falls asleep
[1518-11-14 00:03] falls asleep
[1518-05-14 00:45] falls asleep
[1518-09-20 23:57] Guard #1601 begins shift
[1518-06-27 00:40] falls asleep
[1518-09-20 00:39] falls asleep
[1518-07-17 00:17] falls asleep
[1518-05-09 00:57] falls asleep
[1518-05-20 00:25] falls asleep
[1518-07-06 00:40] wakes up
[1518-04-06 00:13] wakes up
[1518-07-27 00:50] wakes up
[1518-05-06 00:01] Guard #1069 begins shift
[1518-10-11 00:38] wakes up
[1518-06-13 00:57] falls asleep
[1518-05-18 23:57] Guard #1913 begins shift
[1518-06-07 00:50] falls asleep
[1518-04-29 00:50] falls asleep
[1518-07-20 00:15] falls asleep
[1518-08-22 00:52] wakes up
[1518-04-02 00:38] falls asleep
[1518-05-30 00:23] wakes up
[1518-07-02 00:00] falls asleep
[1518-09-17 00:51] wakes up
[1518-10-10 00:42] wakes up
[1518-07-23 00:50] wakes up
[1518-06-04 00:52] wakes up
[1518-06-03 00:45] wakes up
[1518-09-05 00:33] wakes up
[1518-07-16 00:51] wakes up
[1518-06-28 23:59] Guard #2647 begins shift
[1518-07-06 00:18] falls asleep
[1518-11-13 00:11] falls asleep
[1518-08-11 00:03] Guard #881 begins shift
[1518-08-15 00:46] wakes up
[1518-11-07 23:59] Guard #521 begins shift
[1518-06-01 00:58] wakes up
[1518-06-21 00:44] wakes up
[1518-05-17 00:37] wakes up
[1518-04-25 23:57] Guard #1319 begins shift
[1518-08-02 00:18] falls asleep
[1518-04-11 00:49] wakes up
[1518-06-27 00:56] wakes up
[1518-11-04 00:58] wakes up
[1518-06-28 00:54] wakes up
[1518-10-24 00:19] falls asleep
[1518-05-11 00:06] falls asleep
[1518-10-05 00:52] falls asleep
[1518-08-01 00:03] Guard #331 begins shift
[1518-04-10 00:55] wakes up
[1518-06-30 00:44] wakes up
[1518-07-03 00:57] wakes up
[1518-10-06 00:45] falls asleep
[1518-09-23 00:38] falls asleep
[1518-09-11 00:19] falls asleep
[1518-08-08 23:57] Guard #263 begins shift
[1518-04-06 23:56] Guard #521 begins shift
[1518-08-30 00:29] falls asleep
[1518-10-17 00:53] wakes up
[1518-05-09 00:01] Guard #3203 begins shift
[1518-08-10 00:34] falls asleep
[1518-06-16 00:13] falls asleep
[1518-08-02 00:35] wakes up
[1518-11-03 00:49] wakes up
[1518-09-25 00:55] falls asleep
[1518-05-30 00:35] falls asleep
[1518-09-19 00:38] falls asleep
[1518-08-02 00:54] wakes up
[1518-11-12 23:57] Guard #2861 begins shift
[1518-05-28 23:57] Guard #3203 begins shift
[1518-04-07 00:57] wakes up
[1518-07-17 00:02] falls asleep
[1518-08-18 00:30] falls asleep
[1518-06-30 00:19] falls asleep
[1518-09-17 00:48] falls asleep
[1518-07-05 00:35] wakes up
[1518-08-25 00:58] wakes up
[1518-06-07 00:36] falls asleep
[1518-10-27 23:57] Guard #2221 begins shift
[1518-09-09 00:00] Guard #101 begins shift
[1518-11-21 00:37] falls asleep
[1518-07-15 00:08] falls asleep
[1518-05-25 00:30] wakes up
[1518-10-27 00:28] falls asleep
[1518-05-16 00:53] wakes up
[1518-05-06 00:49] falls asleep
[1518-05-12 00:47] falls asleep
[1518-07-27 00:03] Guard #521 begins shift
[1518-06-24 00:56] wakes up
[1518-07-07 00:57] wakes up
[1518-09-24 00:54] wakes up
[1518-09-23 00:02] Guard #2221 begins shift
[1518-08-06 00:31] wakes up
[1518-10-07 00:56] wakes up
[1518-11-04 00:00] Guard #881 begins shift
[1518-05-09 00:06] falls asleep
[1518-09-21 00:32] falls asleep
[1518-09-04 00:36] falls asleep
[1518-10-18 00:37] wakes up
[1518-06-19 00:04] Guard #521 begins shift
[1518-08-29 00:32] wakes up
[1518-05-20 23:58] Guard #2539 begins shift
[1518-09-12 00:04] Guard #1237 begins shift
[1518-07-15 23:58] Guard #1913 begins shift
[1518-05-10 00:59] wakes up
[1518-10-31 00:30] wakes up
[1518-04-22 00:54] wakes up
[1518-09-24 00:50] falls asleep
[1518-07-29 00:58] wakes up
[1518-07-03 00:49] falls asleep
[1518-04-05 00:47] falls asleep
[1518-08-18 00:36] wakes up
[1518-10-03 00:51] wakes up
[1518-08-27 00:46] wakes up
[1518-06-29 00:54] wakes up
[1518-05-25 00:49] wakes up
[1518-04-14 00:34] falls asleep
[1518-07-18 23:53] Guard #1913 begins shift
[1518-06-06 00:08] falls asleep
[1518-06-21 00:57] wakes up
[1518-08-28 00:54] wakes up
[1518-05-01 00:26] falls asleep
[1518-11-14 00:48] wakes up
[1518-08-02 00:00] Guard #521 begins shift
[1518-09-04 00:01] Guard #953 begins shift
[1518-10-16 23:57] Guard #3203 begins shift
[1518-04-23 00:46] falls asleep
[1518-08-02 00:44] falls asleep
[1518-11-11 00:19] falls asleep
[1518-08-20 23:46] Guard #881 begins shift
[1518-05-03 00:00] Guard #101 begins shift
[1518-04-28 00:33] falls asleep
[1518-07-09 00:59] wakes up
[1518-05-11 00:51] wakes up
[1518-10-03 23:52] Guard #2647 begins shift
[1518-09-26 00:57] wakes up
[1518-10-26 00:56] wakes up
[1518-06-08 00:30] falls asleep
[1518-08-23 00:55] wakes up
[1518-11-12 00:39] wakes up
[1518-04-27 00:54] falls asleep
[1518-05-29 00:19] falls asleep
[1518-05-12 00:59] wakes up
[1518-10-12 00:04] Guard #1069 begins shift
[1518-05-15 00:58] wakes up
[1518-06-01 00:37] wakes up
[1518-10-21 00:03] Guard #1601 begins shift
[1518-08-21 00:03] falls asleep
[1518-10-29 00:04] Guard #2861 begins shift
[1518-05-01 00:46] falls asleep
[1518-04-15 00:02] Guard #953 begins shift
[1518-07-11 00:01] falls asleep
[1518-09-11 00:48] wakes up
[1518-07-18 00:15] falls asleep
[1518-07-17 00:11] wakes up
[1518-07-13 00:41] falls asleep
[1518-04-04 00:37] wakes up
[1518-10-14 00:23] falls asleep
[1518-05-07 00:19] falls asleep
[1518-11-17 00:32] falls asleep
[1518-10-12 00:43] wakes up
[1518-11-01 00:28] falls asleep
[1518-11-15 00:01] Guard #1069 begins shift
[1518-07-17 23:56] Guard #1619 begins shift
[1518-04-04 00:58] wakes up
[1518-05-26 00:49] wakes up
[1518-08-15 00:20] wakes up
[1518-04-12 00:37] falls asleep
[1518-08-12 00:33] wakes up
[1518-05-18 00:30] falls asleep
[1518-08-28 00:04] Guard #1319 begins shift
[1518-07-28 23:49] Guard #1319 begins shift
[1518-11-02 00:45] falls asleep
[1518-04-28 00:51] falls asleep
[1518-06-09 00:25] falls asleep
[1518-10-03 00:00] Guard #1619 begins shift
[1518-06-18 00:24] wakes up
[1518-06-14 00:35] falls asleep
[1518-08-04 00:34] wakes up
[1518-06-12 00:42] wakes up
[1518-08-03 00:46] wakes up
[1518-05-05 00:56] wakes up
[1518-11-09 00:23] falls asleep
[1518-05-19 00:47] falls asleep
[1518-05-21 00:37] falls asleep
[1518-06-27 00:16] falls asleep
[1518-06-11 00:07] falls asleep
[1518-07-22 00:52] wakes up
[1518-11-10 00:26] falls asleep
[1518-09-07 00:00] falls asleep
[1518-09-22 00:59] wakes up
[1518-10-31 00:00] Guard #101 begins shift
[1518-09-23 00:33] wakes up
[1518-11-21 00:29] falls asleep
[1518-09-24 00:35] falls asleep
[1518-11-20 00:34] wakes up
[1518-05-01 00:37] wakes up
[1518-06-15 23:59] Guard #2221 begins shift
[1518-06-04 00:00] Guard #1619 begins shift
[1518-05-14 23:56] Guard #421 begins shift
[1518-10-24 00:45] wakes up
[1518-04-06 00:02] Guard #1069 begins shift
[1518-04-10 23:58] Guard #2861 begins shift
[1518-04-30 00:14] wakes up
[1518-06-23 00:39] wakes up
[1518-05-25 00:38] falls asleep
[1518-04-30 00:00] Guard #2539 begins shift
[1518-10-22 00:33] falls asleep
[1518-04-01 00:54] wakes up
[1518-05-11 00:37] falls asleep
[1518-06-04 00:47] falls asleep
[1518-10-24 00:10] wakes up
[1518-08-11 00:52] wakes up
[1518-08-10 00:55] wakes up
[1518-05-03 23:56] Guard #953 begins shift
[1518-04-20 23:56] Guard #3167 begins shift
[1518-11-20 00:02] falls asleep
[1518-04-01 23:59] Guard #881 begins shift
[1518-08-25 00:30] falls asleep
[1518-04-03 00:43] falls asleep
[1518-04-30 00:41] falls asleep
[1518-07-22 00:04] Guard #953 begins shift
[1518-07-27 23:56] Guard #809 begins shift
[1518-11-09 00:57] wakes up
[1518-09-08 00:51] wakes up
[1518-11-06 23:56] Guard #2647 begins shift
[1518-08-25 00:57] falls asleep
[1518-07-25 00:44] falls asleep
[1518-06-13 00:58] wakes up
[1518-10-18 00:59] wakes up
[1518-04-23 00:00] Guard #2861 begins shift
[1518-10-18 00:00] Guard #1319 begins shift
[1518-05-09 00:59] wakes up
[1518-10-30 00:04] falls asleep
[1518-06-26 00:57] wakes up
[1518-09-02 23:56] Guard #421 begins shift
[1518-09-18 00:02] Guard #3203 begins shift
[1518-08-26 00:01] Guard #2861 begins shift
[1518-08-20 00:14] wakes up
[1518-07-26 00:31] falls asleep
[1518-05-30 00:11] falls asleep
[1518-05-06 00:53] wakes up
[1518-05-31 00:30] falls asleep
[1518-09-07 00:46] wakes up
[1518-09-22 00:34] falls asleep
[1518-08-19 00:58] wakes up
[1518-11-02 00:01] falls asleep
[1518-04-14 00:46] falls asleep
[1518-04-13 00:02] falls asleep
[1518-07-31 00:04] Guard #809 begins shift
[1518-07-15 00:40] falls asleep
[1518-04-02 00:25] falls asleep
[1518-11-21 00:52] wakes up
[1518-10-09 23:50] Guard #1237 begins shift
[1518-10-31 00:58] wakes up
[1518-10-25 00:55] falls asleep
[1518-07-15 00:55] wakes up
[1518-11-12 00:26] falls asleep
[1518-07-04 00:46] falls asleep
[1518-06-23 00:36] falls asleep
[1518-10-02 00:03] Guard #809 begins shift
[1518-10-27 00:53] wakes up
[1518-07-29 00:17] wakes up
[1518-10-24 00:30] falls asleep
[1518-04-07 00:45] falls asleep
[1518-10-12 00:37] falls asleep
[1518-05-16 00:51] falls asleep
[1518-06-10 00:21] falls asleep
[1518-09-19 00:18] falls asleep
[1518-08-24 00:01] Guard #1319 begins shift
[1518-06-26 00:37] wakes up
[1518-04-28 00:00] Guard #3203 begins shift
[1518-04-24 00:41] wakes up
[1518-05-23 00:56] wakes up
[1518-11-23 00:52] falls asleep
[1518-08-22 00:36] wakes up
[1518-07-03 00:53] falls asleep
[1518-05-04 00:43] falls asleep
[1518-07-19 00:27] wakes up
[1518-08-06 23:50] Guard #2861 begins shift
[1518-07-25 00:11] falls asleep
[1518-07-12 00:01] falls asleep
[1518-05-20 00:39] falls asleep
[1518-08-24 00:33] falls asleep
[1518-07-06 00:00] Guard #2861 begins shift
[1518-07-07 00:00] falls asleep
[1518-09-18 00:54] wakes up
[1518-04-16 00:22] falls asleep
[1518-09-03 00:06] falls asleep
[1518-08-13 00:21] falls asleep
[1518-10-08 00:33] wakes up
[1518-09-29 00:00] Guard #1601 begins shift
[1518-11-13 23:54] Guard #1601 begins shift
[1518-07-07 23:47] Guard #3203 begins shift
[1518-04-24 00:01] Guard #881 begins shift
[1518-11-07 00:49] falls asleep
[1518-10-08 00:55] wakes up
[1518-09-23 00:19] falls asleep
[1518-08-17 23:59] Guard #2539 begins shift
[1518-06-22 00:32] falls asleep
[1518-04-07 00:09] falls asleep
[1518-07-27 00:57] falls asleep
[1518-06-05 00:03] Guard #89 begins shift
[1518-09-30 00:07] wakes up
[1518-10-19 00:38] falls asleep
[1518-10-23 00:12] wakes up
[1518-09-30 00:04] falls asleep
[1518-06-07 00:42] wakes up
[1518-10-30 00:25] falls asleep
[1518-10-19 00:54] falls asleep
[1518-09-06 23:46] Guard #263 begins shift
[1518-07-12 00:57] wakes up
[1518-08-15 00:04] Guard #2647 begins shift
[1518-10-28 00:51] wakes up
[1518-07-27 00:58] wakes up
[1518-10-01 00:47] falls asleep
[1518-04-11 00:38] falls asleep
[1518-05-17 00:14] falls asleep
[1518-04-05 00:31] wakes up
[1518-10-25 00:58] wakes up
[1518-04-30 23:59] Guard #2861 begins shift
[1518-09-25 23:59] Guard #3203 begins shift
[1518-11-11 00:50] wakes up
[1518-09-12 00:55] falls asleep
[1518-10-13 00:28] wakes up
[1518-08-27 00:18] falls asleep
[1518-07-08 00:57] wakes up
[1518-08-26 00:15] falls asleep
[1518-05-18 00:03] falls asleep
[1518-11-17 00:27] falls asleep
[1518-05-28 00:43] wakes up
[1518-07-03 23:48] Guard #2861 begins shift
[1518-10-04 00:58] wakes up
[1518-05-01 23:57] Guard #2221 begins shift
[1518-04-22 00:36] falls asleep
[1518-05-28 00:54] wakes up
[1518-05-23 00:04] falls asleep
[1518-09-21 00:34] wakes up
[1518-09-19 00:01] falls asleep
[1518-07-01 00:01] Guard #101 begins shift
[1518-05-16 00:31] falls asleep
[1518-11-23 00:44] wakes up
[1518-07-14 00:44] wakes up
[1518-10-20 00:03] Guard #3167 begins shift
[1518-04-27 00:39] wakes up
[1518-07-22 00:09] falls asleep
[1518-08-29 00:00] Guard #1237 begins shift
[1518-07-05 00:53] wakes up
[1518-11-20 00:54] wakes up
[1518-05-11 23:57] Guard #881 begins shift
[1518-05-11 00:00] Guard #521 begins shift
[1518-06-20 00:00] Guard #1601 begins shift
[1518-09-23 00:48] falls asleep
[1518-06-09 00:00] Guard #3167 begins shift
[1518-06-19 00:20] wakes up
[1518-05-02 00:59] wakes up
[1518-10-12 00:58] wakes up
[1518-05-20 00:34] wakes up
[1518-11-05 00:00] Guard #1619 begins shift
[1518-04-07 00:28] wakes up
[1518-04-22 00:44] falls asleep
[1518-07-31 00:47] wakes up
[1518-07-01 00:23] wakes up
[1518-10-01 00:55] wakes up
[1518-09-10 00:58] wakes up
[1518-10-26 00:52] falls asleep
[1518-10-14 00:45] wakes up
[1518-08-25 00:42] falls asleep
[1518-04-15 00:53] wakes up
[1518-10-13 00:02] Guard #1237 begins shift
[1518-05-27 00:20] falls asleep
[1518-07-26 00:19] falls asleep
[1518-07-28 00:59] wakes up
[1518-04-09 00:18] wakes up
[1518-06-11 00:38] falls asleep
[1518-10-05 00:56] wakes up
[1518-11-04 00:42] wakes up
[1518-09-12 00:57] wakes up
[1518-07-30 00:05] falls asleep
[1518-05-03 00:28] falls asleep
[1518-04-14 00:41] wakes up
[1518-10-13 00:27] falls asleep
[1518-05-01 00:58] wakes up
[1518-08-23 00:40] wakes up
[1518-04-18 00:14] falls asleep
[1518-08-11 00:39] falls asleep
[1518-07-14 00:00] Guard #1601 begins shift
[1518-10-30 00:41] wakes up
[1518-06-16 00:46] falls asleep
[1518-08-07 23:54] Guard #1601 begins shift
[1518-10-28 00:30] falls asleep
[1518-04-06 00:50] wakes up
[1518-05-04 00:55] wakes up
[1518-10-14 00:01] Guard #421 begins shift
[1518-11-15 00:55] falls asleep
[1518-11-20 00:53] falls asleep
[1518-05-29 00:47] wakes up
[1518-11-19 00:44] falls asleep
[1518-05-26 00:30] falls asleep
[1518-07-25 23:58] Guard #2441 begins shift
[1518-08-25 00:35] wakes up
[1518-07-06 00:29] wakes up
[1518-10-26 00:43] wakes up
[1518-11-13 00:49] wakes up
[1518-06-10 00:53] wakes up
[1518-05-22 23:53] Guard #2441 begins shift
[1518-09-26 00:39] falls asleep
[1518-07-19 00:00] falls asleep
[1518-08-25 00:52] wakes up
[1518-09-29 00:49] wakes up
[1518-04-26 00:14] falls asleep
[1518-10-05 00:38] wakes up
[1518-07-06 23:47] Guard #421 begins shift
[1518-09-03 00:46] wakes up
[1518-09-28 00:01] Guard #953 begins shift
[1518-09-18 23:49] Guard #1319 begins shift
[1518-10-10 00:02] wakes up
[1518-06-24 00:30] wakes up
[1518-05-15 00:43] falls asleep
[1518-11-19 00:58] wakes up
[1518-08-21 00:54] wakes up
[1518-06-09 00:47] falls asleep
[1518-06-27 23:57] Guard #2647 begins shift
[1518-07-11 23:50] Guard #881 begins shift
[1518-05-24 23:59] Guard #2221 begins shift
[1518-04-08 00:20] falls asleep
[1518-04-30 00:12] falls asleep
[1518-05-21 00:57] wakes up
[1518-10-13 00:57] wakes up
[1518-04-27 00:15] falls asleep
[1518-09-04 00:54] wakes up
[1518-06-25 00:34] falls asleep
[1518-11-18 00:18] falls asleep
[1518-06-14 00:47] wakes up
[1518-09-08 00:55] falls asleep
[1518-07-06 00:37] falls asleep
[1518-07-19 00:55] wakes up
[1518-08-23 00:08] falls asleep
[1518-04-23 00:56] wakes up
[1518-08-15 00:39] falls asleep
[1518-09-11 00:00] Guard #521 begins shift
[1518-10-08 00:03] Guard #881 begins shift
[1518-10-06 00:56] wakes up
[1518-06-03 00:24] falls asleep
[1518-07-12 00:15] wakes up
[1518-10-25 00:00] Guard #2647 begins shift
[1518-06-10 00:51] falls asleep
[1518-11-09 00:39] wakes up
[1518-08-12 00:54] falls asleep
[1518-07-28 00:53] wakes up
[1518-10-17 00:25] falls asleep
[1518-09-25 00:01] Guard #263 begins shift
[1518-06-12 23:59] Guard #101 begins shift
[1518-11-06 00:34] falls asleep
[1518-08-07 00:00] falls asleep
[1518-08-16 23:52] Guard #521 begins shift
[1518-06-06 00:53] wakes up
[1518-07-24 00:04] Guard #421 begins shift
[1518-11-17 00:04] Guard #1913 begins shift
[1518-08-03 00:41] falls asleep
[1518-05-16 00:56] falls asleep
[1518-11-01 00:03] Guard #2221 begins shift
[1518-08-13 00:48] wakes up
[1518-04-19 00:23] falls asleep
[1518-09-25 00:46] wakes up
[1518-06-16 00:54] wakes up
[1518-07-16 00:17] falls asleep
[1518-09-17 00:20] wakes up
[1518-07-24 00:43] falls asleep
[1518-09-14 00:55] wakes up
[1518-04-04 00:51] falls asleep
[1518-08-05 00:04] Guard #89 begins shift
[1518-05-10 00:55] falls asleep
[1518-07-08 00:00] falls asleep
[1518-07-10 00:23] falls asleep
[1518-07-12 00:18] falls asleep
[1518-10-13 00:35] falls asleep
[1518-04-02 00:29] wakes up
[1518-11-16 00:25] falls asleep
[1518-07-18 00:45] wakes up
[1518-05-08 00:01] Guard #1601 begins shift
[1518-08-22 00:28] falls asleep
[1518-04-29 00:04] Guard #1319 begins shift
[1518-06-11 00:59] wakes up
[1518-06-27 00:03] Guard #809 begins shift
[1518-05-11 00:20] wakes up
[1518-07-10 00:03] Guard #421 begins shift
[1518-10-26 00:06] falls asleep
[1518-11-14 00:54] falls asleep
[1518-11-08 00:10] falls asleep
[1518-04-24 00:37] falls asleep
[1518-08-19 00:57] falls asleep
[1518-11-23 00:57] wakes up
[1518-06-20 00:33] falls asleep
[1518-08-29 23:56] Guard #521 begins shift
[1518-11-15 00:57] wakes up
[1518-04-21 23:56] Guard #521 begins shift
[1518-08-26 23:57] Guard #809 begins shift
[1518-04-16 23:57] Guard #1069 begins shift
[1518-10-26 23:56] Guard #2861 begins shift
[1518-11-17 00:57] wakes up
[1518-11-20 23:56] Guard #1601 begins shift
[1518-11-18 23:59] Guard #3203 begins shift
[1518-05-31 00:02] Guard #1319 begins shift
[1518-07-08 23:46] Guard #953 begins shift
[1518-06-04 00:35] falls asleep
[1518-08-31 00:32] falls asleep
[1518-04-03 00:08] falls asleep
[1518-07-29 00:05] falls asleep
[1518-07-20 00:55] wakes up
[1518-05-16 00:58] wakes up
[1518-06-21 00:18] falls asleep
[1518-11-21 00:34] wakes up
[1518-05-20 00:53] wakes up
[1518-07-23 00:28] falls asleep
[1518-08-19 00:47] wakes up
[1518-04-06 00:24] falls asleep
[1518-11-01 00:48] wakes up
[1518-10-21 00:24] falls asleep
[1518-06-06 00:26] wakes up
[1518-07-05 00:15] falls asleep
[1518-10-02 00:45] wakes up
[1518-08-15 00:19] falls asleep
[1518-04-10 00:00] Guard #881 begins shift
[1518-10-29 23:51] Guard #3167 begins shift
[1518-11-03 00:27] wakes up
[1518-09-08 00:00] Guard #3203 begins shift
[1518-09-17 00:41] falls asleep
[1518-05-31 00:44] falls asleep
[1518-10-15 00:27] falls asleep
[1518-10-06 00:00] Guard #1319 begins shift
[1518-04-14 00:47] wakes up
[1518-07-25 00:26] wakes up
[1518-09-17 00:45] wakes up
[1518-04-30 00:58] wakes up
[1518-05-22 00:15] falls asleep
[1518-04-19 00:48] wakes up
[1518-07-26 00:28] wakes up
[1518-07-31 00:41] falls asleep
[1518-07-17 00:39] wakes up
[1518-07-01 00:28] falls asleep
[1518-10-16 00:01] Guard #1237 begins shift
[1518-08-29 00:30] falls asleep
[1518-04-13 23:59] Guard #2861 begins shift
[1518-07-10 00:28] wakes up
[1518-10-24 00:04] falls asleep
[1518-07-28 00:48] wakes up
[1518-11-10 23:57] Guard #1619 begins shift
[1518-11-09 00:44] falls asleep
[1518-10-10 00:01] falls asleep
[1518-05-06 00:58] wakes up
[1518-09-22 00:54] falls asleep
[1518-09-05 00:54] wakes up
[1518-07-29 23:48] Guard #1619 begins shift
[1518-10-11 00:47] falls asleep
[1518-07-12 00:54] falls asleep
[1518-05-11 00:54] falls asleep
[1518-07-26 00:36] wakes up
[1518-07-28 00:47] falls asleep
[1518-05-31 00:16] wakes up
[1518-09-08 00:11] falls asleep
[1518-06-23 23:58] Guard #263 begins shift
[1518-10-09 00:02] falls asleep
[1518-11-22 00:57] wakes up
[1518-05-08 00:08] falls asleep
[1518-10-23 00:11] falls asleep
[1518-10-09 00:34] wakes up
[1518-08-14 00:04] Guard #2539 begins shift
[1518-04-28 00:54] wakes up
[1518-06-30 00:00] Guard #1069 begins shift
[1518-06-28 00:47] wakes up
[1518-10-15 00:00] Guard #3167 begins shift
[1518-06-25 23:56] Guard #2221 begins shift
[1518-07-18 00:23] wakes up
[1518-08-08 00:49] wakes up
[1518-09-27 00:06] falls asleep
[1518-07-05 00:44] wakes up
[1518-06-11 00:31] wakes up
[1518-07-29 00:50] falls asleep
[1518-09-27 00:02] Guard #2647 begins shift
[1518-07-18 00:30] falls asleep
[1518-06-28 00:52] falls asleep
[1518-06-10 00:24] wakes up
[1518-08-22 00:00] Guard #1601 begins shift
[1518-09-28 00:37] falls asleep
[1518-07-04 00:48] wakes up
[1518-09-15 00:06] falls asleep
[1518-05-04 00:21] falls asleep
[1518-08-14 00:44] wakes up
[1518-10-20 00:35] wakes up
[1518-05-17 23:51] Guard #2647 begins shift
[1518-07-23 00:02] Guard #1619 begins shift
[1518-09-10 00:40] falls asleep
[1518-10-03 00:57] wakes up
[1518-11-16 00:03] Guard #1619 begins shift
[1518-06-22 00:43] falls asleep
[1518-06-22 00:58] wakes up
[1518-09-11 00:39] wakes up
[1518-09-20 00:02] Guard #1619 begins shift
[1518-10-04 00:19] wakes up
[1518-10-31 00:34] falls asleep
[1518-05-14 00:39] wakes up
[1518-11-01 23:49] Guard #1601 begins shift
[1518-04-03 00:00] Guard #881 begins shift
[1518-05-05 00:32] falls asleep
[1518-04-12 00:02] Guard #1619 begins shift
[1518-05-26 00:01] Guard #1913 begins shift
[1518-07-09 00:04] falls asleep
[1518-09-18 00:30] falls asleep
[1518-08-20 00:12] falls asleep
[1518-09-19 00:27] wakes up
[1518-07-15 00:02] Guard #3203 begins shift
[1518-04-08 00:55] wakes up
[1518-04-09 00:00] Guard #881 begins shift
[1518-04-15 00:12] falls asleep
[1518-10-20 00:26] falls asleep
[1518-08-19 00:42] falls asleep
[1518-11-16 00:48] wakes up
[1518-04-16 00:37] wakes up
[1518-05-14 00:27] falls asleep
[1518-09-29 00:23] falls asleep
[1518-11-06 00:27] wakes up
[1518-09-10 00:51] wakes up
[1518-08-09 00:08] falls asleep
[1518-11-10 00:52] wakes up
[1518-08-24 23:57] Guard #1601 begins shift
[1518-11-22 00:45] falls asleep
[1518-11-02 00:23] wakes up
[1518-09-16 00:57] wakes up
[1518-11-06 00:00] Guard #2861 begins shift
[1518-05-18 00:26] wakes up
[1518-08-10 00:26] wakes up
[1518-09-28 00:54] wakes up
[1518-06-01 00:12] falls asleep
[1518-10-23 00:19] falls asleep
[1518-09-05 00:00] Guard #1601 begins shift
[1518-09-11 00:43] falls asleep
[1518-07-28 00:57] falls asleep
[1518-05-10 00:32] wakes up
[1518-10-05 00:04] Guard #881 begins shift
[1518-04-29 00:58] wakes up
[1518-11-04 00:22] falls asleep
[1518-09-25 00:15] falls asleep
[1518-10-23 00:01] Guard #3167 begins shift
[1518-09-30 23:56] Guard #2647 begins shift
[1518-08-10 00:03] falls asleep
[1518-06-20 00:54] wakes up
[1518-10-22 00:02] falls asleep
[1518-05-03 00:48] wakes up
[1518-07-28 00:51] falls asleep
[1518-08-28 00:23] falls asleep
[1518-11-14 00:56] wakes up
[1518-10-21 00:09] falls asleep
[1518-08-31 00:24] falls asleep
[1518-05-16 00:36] wakes up
[1518-04-08 00:00] Guard #521 begins shift
[1518-10-12 00:53] falls asleep
[1518-09-12 00:20] falls asleep
[1518-11-21 23:58] Guard #2221 begins shift
[1518-05-22 00:43] wakes up
[1518-10-22 00:51] wakes up
[1518-04-01 00:53] falls asleep
[1518-06-27 00:51] falls asleep
[1518-07-19 00:48] falls asleep
[1518-04-22 00:30] wakes up
[1518-05-07 00:51] wakes up
[1518-09-21 23:57] Guard #2221 begins shift
[1518-09-09 00:31] falls asleep
[1518-05-04 00:40] wakes up
[1518-06-28 00:32] falls asleep
[1518-08-06 00:19] falls asleep
[1518-05-24 00:00] Guard #421 begins shift
[1518-10-16 00:36] wakes up
[1518-06-24 00:51] falls asleep
[1518-05-09 00:54] wakes up
[1518-05-21 00:54] falls asleep
[1518-04-29 00:18] falls asleep
[1518-06-21 00:49] falls asleep
[1518-04-15 00:13] wakes up
[1518-08-11 00:55] falls asleep
[1518-10-29 00:55] wakes up
[1518-07-09 00:35] wakes up
[1518-07-24 00:53] wakes up
[1518-06-26 00:46] falls asleep
[1518-04-09 00:11] falls asleep
[1518-05-27 00:52] wakes up
[1518-04-02 00:54] wakes up
[1518-10-18 00:46] wakes up
[1518-07-04 00:05] falls asleep
[1518-04-17 00:53] wakes up
[1518-11-03 00:48] falls asleep
[1518-11-18 00:53] wakes up
[1518-05-20 00:03] Guard #1619 begins shift
[1518-07-25 00:53] wakes up
[1518-08-22 00:40] falls asleep
[1518-08-14 00:42] falls asleep
[1518-10-21 00:51] wakes up
[1518-06-20 00:38] wakes up
[1518-08-31 00:27] wakes up
[1518-11-19 23:54] Guard #2441 begins shift
[1518-10-08 00:41] falls asleep
[1518-07-06 00:43] falls asleep
[1518-09-23 00:43] wakes up
[1518-09-06 00:59] wakes up
[1518-09-13 00:51] wakes up
[1518-10-22 00:25] wakes up
[1518-09-13 00:36] falls asleep
[1518-06-22 00:00] Guard #2221 begins shift
[1518-04-03 00:18] wakes up
[1518-05-02 00:52] falls asleep
[1518-06-25 00:46] wakes up
[1518-08-22 23:59] Guard #809 begins shift
[1518-08-30 00:42] wakes up
[1518-08-12 00:58] wakes up
[1518-04-12 00:58] wakes up
[1518-08-24 00:54] wakes up
[1518-05-14 00:02] Guard #1069 begins shift
[1518-05-21 00:43] falls asleep
[1518-11-05 00:18] falls asleep
[1518-07-17 00:45] wakes up
[1518-06-09 00:31] wakes up
[1518-07-11 00:56] wakes up
[1518-04-13 00:56] wakes up
[1518-08-04 00:09] falls asleep
[1518-08-18 23:57] Guard #2647 begins shift
[1518-07-02 00:46] wakes up
[1518-10-06 23:59] Guard #521 begins shift
[1518-08-26 00:53] wakes up
[1518-09-12 00:46] wakes up
[1518-06-14 23:57] Guard #1913 begins shift
[1518-05-08 00:44] wakes up
[1518-06-23 00:01] Guard #2539 begins shift
[1518-04-17 00:51] falls asleep
[1518-11-17 00:29] wakes up
[1518-09-14 00:27] falls asleep
[1518-07-14 00:18] falls asleep
[1518-09-25 00:56] wakes up
[1518-05-12 23:59] Guard #521 begins shift
[1518-07-01 00:34] wakes up
[1518-08-31 00:45] wakes up
[1518-08-12 00:02] Guard #1601 begins shift
[1518-11-03 00:21] falls asleep
[1518-05-31 00:07] falls asleep
[1518-05-05 00:01] Guard #1601 begins shift
[1518-05-10 00:01] Guard #1069 begins shift
[1518-07-20 23:56] Guard #89 begins shift
[1518-09-23 00:53] wakes up
[1518-07-16 23:50] Guard #2539 begins shift
[1518-04-04 23:56] Guard #809 begins shift
[1518-09-24 00:44] wakes up
[1518-05-11 00:56] wakes up
[1518-09-20 00:53] wakes up
[1518-05-19 00:37] falls asleep
[1518-06-26 00:09] falls asleep
[1518-04-24 00:54] wakes up
[1518-07-12 23:59] Guard #2861 begins shift
[1518-10-04 00:43] falls asleep
[1518-05-29 00:53] falls asleep
[1518-05-19 00:39] wakes up
[1518-05-19 00:48] wakes up
[1518-05-26 23:59] Guard #2539 begins shift
[1518-06-11 23:59] Guard #421 begins shift
[1518-10-18 00:42] falls asleep
[1518-07-04 23:56] Guard #2441 begins shift
[1518-07-13 00:48] wakes up
[1518-06-10 00:02] Guard #263 begins shift
[1518-09-10 00:57] falls asleep
[1518-04-21 00:26] falls asleep
[1518-10-16 00:23] falls asleep
[1518-06-27 00:28] wakes up
[1518-11-04 00:51] falls asleep
[1518-08-12 00:28] falls asleep
[1518-11-21 00:38] wakes up
[1518-04-12 23:54] Guard #1319 begins shift
[1518-06-15 00:43] falls asleep
[1518-04-24 00:46] falls asleep
[1518-05-06 00:57] falls asleep
[1518-07-27 00:27] falls asleep
[1518-06-09 00:58] wakes up
[1518-10-10 00:16] falls asleep
[1518-06-21 00:00] Guard #953 begins shift
[1518-05-24 00:52] wakes up
[1518-04-16 00:01] Guard #1069 begins shift
[1518-10-19 00:00] Guard #881 begins shift
[1518-10-15 00:30] wakes up
[1518-05-31 23:59] Guard #2221 begins shift
[1518-09-22 00:50] wakes up
[1518-09-17 00:02] Guard #1619 begins shift
[1518-08-20 00:49] wakes up
[1518-10-21 23:46] Guard #521 begins shift
[1518-06-01 23:57] Guard #1319 begins shift
[1518-10-18 00:30] falls asleep
[1518-05-21 23:56] Guard #1619 begins shift
[1518-09-04 00:44] falls asleep
[1518-04-05 00:52] wakes up
[1518-09-15 23:59] Guard #881 begins shift
[1518-10-21 00:19] wakes up
[1518-06-07 00:54] wakes up
[1518-07-06 00:58] wakes up
[1518-04-04 00:00] Guard #809 begins shift
[1518-06-18 00:09] falls asleep
[1518-09-19 00:55] wakes up
[1518-09-13 23:56] Guard #521 begins shift
[1518-11-12 00:01] Guard #101 begins shift
[1518-09-06 00:31] falls asleep
[1518-06-24 00:11] falls asleep
[1518-09-14 23:56] Guard #2539 begins shift
[1518-06-04 00:44] wakes up
[1518-05-30 00:04] Guard #3203 begins shift
[1518-05-30 00:56] wakes up
[1518-09-16 00:13] falls asleep
[1518-05-08 00:49] falls asleep
[1518-11-10 00:01] Guard #1619 begins shift
[1518-11-05 00:51] wakes up
[1518-05-15 00:48] wakes up
[1518-07-02 23:56] Guard #1619 begins shift
[1518-11-18 00:02] Guard #3203 begins shift
[1518-04-05 00:27] falls asleep
[1518-08-23 00:51] falls asleep
[1518-10-18 00:49] falls asleep
[1518-11-06 00:52] wakes up
[1518-08-06 00:01] Guard #2441 begins shift
[1518-05-13 00:59] wakes up
[1518-04-28 00:37] wakes up
[1518-06-19 00:18] falls asleep
[1518-09-19 00:12] wakes up
[1518-05-08 00:53] wakes up
[1518-04-22 00:40] wakes up
[1518-11-08 00:11] wakes up
[1518-07-03 00:50] wakes up
[1518-10-19 00:59] wakes up
[1518-09-03 00:49] falls asleep
[1518-11-08 00:38] wakes up
[1518-08-03 00:02] Guard #2861 begins shift
[1518-04-22 00:29] falls asleep
[1518-10-08 23:47] Guard #1601 begins shift
[1518-10-31 00:29] falls asleep
[1518-10-03 00:16] falls asleep
[1518-07-17 00:42] falls asleep
[1518-04-20 00:42] wakes up
[1518-04-04 00:30] falls asleep
[1518-05-31 00:40] wakes up
[1518-08-17 00:47] wakes up
[1518-05-18 00:57] wakes up`

type guard struct {
	id     int
	shifts []shift
}

func (g *guard) totalSleep() int {
	var t int
	for _, s := range g.shifts {
		for _, n := range s.naps {
			t += n.wake - n.asleep
		}
	}
	return t
}

func (g *guard) maxMinute() (min, n int) {
	freq := make(map[int]int)
	for _, s := range g.shifts {
		for _, n := range s.naps {
			for i := n.asleep; i < n.wake; i++ {
				freq[i]++
			}
		}
	}
	max := 0
	maxMin := 0
	for m, n := range freq {
		if n > max {
			max = n
			maxMin = m
		}
	}
	return maxMin, max
}

type shift struct {
	starthour, startmin int
	naps                []nap
}

type nap struct {
	asleep, wake int
}

func main() {
	// part 1
	guards := make(map[int]guard)
	lines := utils.Lines(input)
	sort.Strings(lines)
	for i := 0; i < len(lines); {
		var id, month, day, starthour, startmin int
		utils.Sscanf(lines[i], "[1518-%d-%d %d:%d] Guard #%d begins shift", &month, &day, &starthour, &startmin, &id)

		var naps []nap
		for i++; i < len(lines) && !strings.Contains(lines[i], "begins"); i += 2 {
			var asleephour, asleepmin, wakehour, wakemin int
			utils.Sscanf(lines[i], "[1518-%d-%d %d:%d] falls asleep", &month, &day, &asleephour, &asleepmin)
			utils.Sscanf(lines[i+1], "[1518-%d-%d %d:%d] wakes up", &month, &day, &wakehour, &wakemin)
			naps = append(naps, nap{asleepmin, wakemin})
		}
		g := guards[id]
		g.id = id
		g.shifts = append(g.shifts, shift{starthour, startmin, naps})
		guards[id] = g
	}
	var max, maxID int
	for _, g := range guards {
		if t := g.totalSleep(); t > max {
			max = t
			maxID = g.id
		}
	}
	g := guards[maxID]
	min, _ := g.maxMinute()
	utils.Println(g.id * min)

	// part 2
	max = 0
	maxID = 0
	for _, g := range guards {
		_, n := g.maxMinute()
		if n > max {
			max = n
			maxID = g.id
		}
	}
	g = guards[maxID]
	min, _ = g.maxMinute()
	utils.Println(min * maxID)
}
