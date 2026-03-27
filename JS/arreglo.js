let arreglo = [1,2,3,4,5,6,7,8,9,10]

function recorrerFor(){
    for ( let i = 0; i < arreglo.length ; i++){
        console.log(arreglo[i])

    }

}

function recorrerForEach(){
    arreglo.forEach(e=>(console.log(e)))
}

function impares (){
    for ( let i = 0; i < arreglo.length ; i++){
        if (arreglo[i] % 2 != 0){
            arreglo[i] = 0
        }
    }

    recorrerFor()
}

function multIndice(){
    for ( let i = 0; i < arreglo.length ; i++){
        arreglo[i] = arreglo[i] * i
    }

    recorrerFor()
}

function busqueda(){
    let valor =3
    for ( let i = 0; i < arreglo.length ; i++){
        if (arreglo[i] == valor){
            console.log("El valor fue encontrado en el arreglo")
        }else {
            console.log("El valor no se encuentra en el arreglo")
        }
    }
}

recorrerFor()
recorrerForEach()
impares()
multIndice()
busqueda()