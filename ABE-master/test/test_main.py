'''
:Authors:         Shashank Agrawal
:Date:            5/2016
'''

from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE


def main():
    # instantiate a bilinear pairing map
    pairing_group = PairingGroup('MNT224')
    print("选线并产生配对群")
    
    # AC17 CP-ABE under DLIN (2-linear)
    cpabe = AC17CPABE(pairing_group, 2)
    print("基于配对群产生cpabe方案")
    
    # run the set up
    (pk, msk) = cpabe.setup()
    print("setup产生主公私钥对",pk,msk,type(pk),type(msk))
    
    # generate a key
    attr_list = ['ONE', 'TWO', 'THREE']
    key = cpabe.keygen(pk, msk, attr_list)
    print("产生了密钥",key,type(key))
    
    # choose a random message
    msg = pairing_group.random(GT)
    print("产生了随机群元素作为消息",msg,type(msg))
    
    # generate a ciphertext
    policy_str = '((ONE and THREE) and (TWO OR FOUR))'
    ctxt = cpabe.encrypt(pk, msg, policy_str)
    print("产生了密文",ctxt,type(ctxt))

    # decryption
    rec_msg = cpabe.decrypt(pk, ctxt, key)
    print("解密出了明文",rec_msg,type(rec_msg))
    
    if debug:
        if rec_msg == msg:
            print ("Successful decryption.")
        else:
            print ("Decryption failed.")


if __name__ == "__main__":
    debug = False
    main()