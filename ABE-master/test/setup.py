from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE
    
    
pairing_group = PairingGroup('MNT224')

# AC17 CP-ABE under DLIN (2-linear)
cpabe = AC17CPABE(pairing_group, 2)

# run the set up
(pk, msk) = cpabe.setup()

# pairing_group,cpabe,pk,msk这四个都是要给别人用的