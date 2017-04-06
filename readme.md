Para gerar o executável
    go build
e depois servir em alguma porta
    PORT=5000 ./fiscaluno-ws
rodar - godep save - para salvar as dependências (usado pelo heroku)

Prestar atenção no gopath, tive vários problemas com ele.
Executar go install pacoteOuDependencia para adicionar a dependencia ao SEU gopath

O nome do arquivo.go deve ser o mesmo que o da última pasta em que ele está contido, assim como o seu pacote

Ex: Olhar o institutionservice

