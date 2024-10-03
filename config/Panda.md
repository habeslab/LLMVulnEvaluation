### Installation
1. Clone the repository, give the permission to the installation file and execute the script.
	```bash
	git clone https://github.com/Sun-C0ffee/Panda
	chmod 777 install.sh
	./install.sh
	```
2. Install the py-algorand-sdk.
	```bash
	pip install py-algorand-sdk==1.20.2
	```
3. Start the tool with the given option.
	```bash
	python3 ./panda.py -sc -s <pathToSmartContract>/test.teal -rs rule1
	```
	