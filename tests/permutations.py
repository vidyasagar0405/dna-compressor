from itertools import permutations
import subprocess


def possible_byte_permutation(allowed: str) -> str:
    perm: list[str] = []
    for p in permutations(allowed):
        perm.append("".join(p))

    return "\n".join(perm)


def main():
    out_file = "./dna-permutations.txt"

    allowed: str = "atgc"
    perms: str = possible_byte_permutation(allowed)

    with open(out_file, "w") as of:
        of.write(perms)

    cmd = f"../dna-compressor {out_file}"
    subprocess.run(cmd.split())

    bin_out = f"{out_file}.encoded"
    with open(bin_out, "rb") as f:
        data = f.read()

    for byte in data:
        print(format(byte, "08b"))


if __name__ == "__main__":
    main()

    """

    atgc
    atcg
    agtc
    agct
    actg
    acgt
    tagc
    tacg
    tgac
    tgca
    tcag
    tcga
    gatc
    gact
    gtac
    gtca
    gcat
    gcta
    catg
    cagt
    ctag
    ctga
    cgat
    cgta

    00011011
    00011110
    00100111
    00101101
    00110110
    00111001
    01001011
    01001110
    01100011
    01101100
    01110010
    01111000
    10000111
    10001101
    10010011
    10011100
    10110001
    10110100
    11000110
    11001001
    11010010
    11011000
    11100001
    11100100

    from golang:
    27
    30
    39
    45
    54
    57
    75
    78
    99
    108
    114
    120
    135
    141
    147
    156
    177
    180
    198
    201
    210
    216
    225
    228

    """
