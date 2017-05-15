# Fiscaluno-ws
Repositório para a criação do webservice do projeto Fiscaluno.
Documentação: https://goo.gl/BU3Rbw

# Antes de começar...

## Crie um arquivo conf.json e adicione-o ao ``.gitignore``
   Por segurança, não há referência sobre a url do firebase, portanto, criar um arquivo chamado ``conf.json`` com a url do projeto. 
 
   Obs: Adicionar o arquivo `conf.json` e adicionar no ``.gitignore``. Não utilizar os comandos ``git add .`` nem ``git commit -a``

### Para gerar o executável
   Execute o comando `` go build``
### e depois sirva em alguma porta
   `` PORT=5000 ./fiscaluno-ws``

## Dependências
   Executar o comando ``godep save`` antes de subir no heroku.

### links úteis:
- http://ernestmicklei.com/2012/11/go-restful-first-working-example/
- http://labix.org/mgo
- https://golang.org/doc/code.html
- https://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html
