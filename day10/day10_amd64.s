// no bounds checking or calls to morestack
// some help optimizing this would be nice...

// func run(b []byte, c byte) int
TEXT ·run(SB), $0-40
start:
    MOVQ    (TLS), CX
    MOVQ    b+0(FP), DI
    MOVBQZX c+24(FP), SI
    MOVQ    b_len+8(FP), DX
    MOVQ    $0, AX
    CMPQ    AX, DX
    JGE     $0, done
loop:
    LEAQ    (DI)(AX*1), BX
    MOVBQZX (BX), BX
    CMPB    BL, SIB
    JNE     $0, done
    INCQ    AX
    CMPQ    AX, DX
    JLT     $0, loop
done:
    MOVQ    AX, ret+32(FP)
    RET
