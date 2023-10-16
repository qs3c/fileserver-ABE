from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE
from setup import cpabe,pk,msk

# 获取属性列表
attr_list = ['ONE', 'TWO', 'THREE']

#产生密钥
key = cpabe.keygen(pk, msk, attr_list)

#输出ABE密钥文件
   