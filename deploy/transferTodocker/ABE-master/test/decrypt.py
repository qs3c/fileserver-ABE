from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE
from setup import cpabe,pk

#读取加密密钥文件
key = ""
#读取用户私钥
ctxt = ""
# decryption
rec_msg = cpabe.decrypt(pk, ctxt, key)

# 截取一定byte长度，然后输出为对称密钥文件！