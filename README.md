# Massdrop Keyboard Utilities

## ledgen
Tool to create the `led_instruction`s from an easier to create TOML file.

 ```toml
[Colors]
A = "FF0000" # Alpha, Space Keys
B = "00FF00" # Punctuation, Symbol Keyslion
C = "0000FF" # Modifier Keys
D = "00FF00" # Navigation Keys
E = "FFFF00" # Number Keys

[[Layers]]
# If provided, MatchLayers will specify which layers this lighting configuration is displayed on.
# If not provided, this lighting configuration will be displayed on all layers and overridden where specified.
# MatchLayers = [0]

KeyLEDCount = 87
KeyLEDs = '''
	A AAAA AAAA AAAA AAA

	AAAAAAAAAAAAAA AAA
	A AAAAAAAAAAAAA AAA
	A AAAAAAAAAAA A
	A AAAAAAAAAA  A   A
	AAA   A   AAAA   AAA
'''

EdgeLEDCount = 32
EdgeLEDs = '''
	DDDDDDDDDDDDD
	DDD
	DDDDDDDDDDDDD
	DDD
'''
 
```
