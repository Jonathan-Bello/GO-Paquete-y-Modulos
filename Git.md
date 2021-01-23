# Subida de un repositorio local a un repositorio remoto ya existente

Para ello lo primero que debemos hacer es realizar commit de nuestros archivos locales.

Para despues ir a crear un repositorio con el nombre que deseemos.

Y existe una sección que nos indica que comando usar, el cual es

    git remote add origin link-github de respositorio
    git remote add origin git@github.com:Jonathan-Bello/GO-Paquete-y-Modulos.git

De esta manera agregamos un origin remoto, lo podemos verficiar con

    git remote -v

Y por ultimo debemos empujar nuestros commits al origen

    git push -u origin master
