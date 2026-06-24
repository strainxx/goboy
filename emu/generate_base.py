import requests
import dataclasses

res = requests.get("https://gbdev.io/gb-opcodes/Opcodes.json").json()

@dataclasses.dataclass
class InstrInfo:
    opcode: str
    mnemonic: str
    cycles: int
    impl_name: str

r8 = ["A", "B", "C", "D", "E", "H", "L"]
r16 = ["AF", "BC", "DE", "HL", "SP"]
ims = ["n8","n16", "e8", "cc", "vec", "a16", "Z", "C", "NZ", "NC", "a8"]
u3 = ["0", "1", "2", "3", "4", "5", "6", "7"]
protos = set()

def gen_full_mnemonic(instr: dict):
    mnem = instr["mnemonic"]

    ops = ""

    for opp in instr["operands"]:
        is_dec = opp.get("decrement", False)
        is_inc = opp.get("increment", False)
        is_imm = opp["immediate"]
        pr = "[" if not is_imm else ""
        sf = "+" if is_inc else "-" if is_dec else ""
        if not is_imm: sf+="]"
        ops += f"{pr}{opp['name']}{sf}, "

    return f"{mnem} {ops}".removesuffix(", ")

# with open("genimpl.go", "w") as f:
#     f.write("package cpu\n\n")
#     for op, instr in res["unprefixed"].items():
#         mnemonic: str = instr['mnemonic']
#         pair = []
#         args = ["cpu"]
#         prot_args = ["cpu *CPU"]
#         for i, opp in enumerate(instr["operands"]):
#             is_dec = opp.get("decrement", False)
#             is_inc = opp.get("increment", False)
#             suf = "inc" if is_inc else 'dec' if is_dec else 'mem' if not opp["immediate"] else ''
#             if opp["name"] in r8:
#                 pair.append("r8")
#                 args.append("&cpu."+opp["name"])
#                 prot_args.append(f"op{i} *byte")
#             elif opp["name"] in r16:
#                 pair.append("r16"+suf)
#                 args.append(opp["name"])
#                 prot_args.append(f"op{i} R16Enum")
#             elif opp["name"] in ims:
#                 pair.append(opp["name"])
#             elif opp["name"].startswith("$"):
#                 pair.append("vec")
#             else:
#                 raise Exception(f"Unknown op {opp['name']} ({op})")
        
#         str_args = ','.join(args) if len(args) != 0 else ''
#         str_prot_args = ','.join(prot_args) if len(prot_args) != 0 else ''
#         gen_name = "Impl_"+op#'_'.join([mnemonic, *[op["name"] for op in instr["operands"]]])
#         # gen_name = gen_name.replace("$", "d")
#         f.write(f"func {gen_name}(cpu *CPU) int {{\n")
#         # f.write(f"\tcpu.Halt(\"{gen_name} not implemented!\")\n")
#         callfunc = ""
#         if len(pair) != 0:
#             callfunc = f"{mnemonic}_{'_'.join(pair)}"
#         else:
#             callfunc = mnemonic
#         f.write(f"\treturn {callfunc}({str_args})\n")
#         # f.write(f"\treturn 0\n")
#         f.write("}\n\n")

#         if False and callfunc not in protos:
#             print(f"""\nfunc {callfunc}({str_prot_args}) int {{
#     cpu.Halt("{callfunc} not implemented!")
#     return 0
# }}\n""")
#             protos.add(callfunc)

#         if True: print(f"""
# Opcodes[{op}] = Instruction{{
# 	Name:   "{gen_full_mnemonic(instr)}",
# 	Bytes:   {instr["bytes"]},
# 	Exec:   {gen_name} ,
# }}
# """)
	# Cycles: {instr['cycles'][0]},
        

# with open("cbimpl.go", "w") as f:
# f.write("package cpu\n\n")
for op, instr in res["cbprefixed"].items():
    mnemonic: str = instr['mnemonic']
    pair = []
    args = ["cpu"]
    prot_args = ["cpu *CPU"]
    for i, opp in enumerate(instr["operands"]):
        is_dec = opp.get("decrement", False)
        is_inc = opp.get("increment", False)
        suf = "inc" if is_inc else 'dec' if is_dec else 'mem' if not opp["immediate"] else ''
        if opp["name"] in r8:
            pair.append("r8")
            args.append("&cpu."+opp["name"])
            prot_args.append(f"op{i} *byte")
        elif opp["name"] in r16:
            pair.append("r16"+suf)
            args.append(opp["name"])
            prot_args.append(f"op{i} R16Enum")
        elif opp["name"] in u3:
            pair.append("u3")
            args.append(opp["name"])
            prot_args.append(f"u3 uint8")
        elif opp["name"] in ims:
            pair.append(opp["name"])
        
        elif opp["name"].startswith("$"):
            pair.append("vec")
        else:
            raise Exception(f"Unknown op {opp['name']} ({op})")
    
    str_args = ','.join(args) if len(args) != 0 else ''
    str_prot_args = ','.join(prot_args) if len(prot_args) != 0 else ''
    gen_name = "Impl_"+op.replace("0x", "0xCB")#'_'.join([mnemonic, *[op["name"] for op in instr["operands"]]])
    # gen_name = gen_name.replace("$", "d")
    # f.write(f"func {gen_name}(cpu *CPU) int {{\n")
    # f.write(f"\tcpu.Halt(\"{gen_name} not implemented!\")\n")
    callfunc = ""
    if len(pair) != 0:
        callfunc = f"{mnemonic}_{'_'.join(pair)}"
    else:
        callfunc = mnemonic
    # f.write(f"\treturn {callfunc}({str_args})\n")
    # f.write(f"\treturn 0\n")
    # f.write("}\n\n")

    if False and callfunc not in protos:
        print(f"""\nfunc {callfunc}({str_prot_args}) int {{
cpu.Halt("{callfunc} not implemented!")
return 0
}}\n""")
        protos.add(callfunc)

    if True: print(f"""
CBOpcodes[{op}] = Instruction{{
	Name:   "{gen_full_mnemonic(instr)}",
 	Bytes:   {instr["bytes"]},
	Exec:   {gen_name} ,
}}
""")