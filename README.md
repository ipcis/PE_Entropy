# PE_Entropy

The Entropy Calculator is a simple command-line tool written in Go that calculates the entropy of data within a file, such as a binary executable, a text file, or any other type of file. Entropy is a measure of randomness or unpredictability within a set of data, and it can be used for various purposes, including file analysis and security assessments.

```
git clone https://github.com/ipcis/PE_Entropy.git
```

```
go build -o pe_entropy.exe pe_entropy.go
```
```
pe_entropy.exe c:\Windows\System32\calc.exe
Section: .text
Entropy: 5.806535
Section: .rdata
Entropy: 3.969362
Section: .data
Entropy: 0.378703
Section: .pdata
Entropy: 1.977328
Section: .rsrc
Entropy: 2.813526
Section: .reloc
Entropy: 0.463472
```
