This repository contains the replication package for the ICST 2025 Submission. In particular, this repository is composed by the following folders:
- config: contains all the information needed to setup the tools to perform the analysis.
- dataset: contains all the smart contracts used for the validation of the tools.
	* Panda
	* VRust
	* HFCCT
In particular, the VRust folder contains two files plus an optional file:
- `changes.pdf`: the changes applied over the original programs.
- `result.pdf`: the result of the analysis performed by the tool.
- `report.png`: the vulnerability found by the tool (if any).

To replicate the analysis performed by our research, you have to install the individual tool and then process the smart contracts contained in the dataset folders. These smart contracts have been generated using ChatGPT4.