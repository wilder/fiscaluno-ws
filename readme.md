## Para gerar o executável
   > go build
## e depois servir em alguma porta
   > PORT=5000 ./fiscaluno-ws
## rodar para salvar as dependências (usado pelo heroku)
   > godep save

**Prestar atenção no gopath, tive vários problemas com ele.**

## Executar para adicionar a dependencia ao SEU gopath
   > go install *pacoteOuDependencia* 

**O nome do arquivo.go deve ser o mesmo que o da última pasta em que ele está contido, assim como o seu pacote**

Ex: Olhar o institutionservice

### links úteis:
- http://ernestmicklei.com/2012/11/go-restful-first-working-example/
- http://labix.org/mgo
- https://golang.org/doc/code.html
- https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
- http://stackoverflow.com/questions/15049903/how-to-use-custom-packages-in-golang
- http://stackoverflow.com/questions/13214029/go-build-cannot-find-package-even-though-gopath-is-set

