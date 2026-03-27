let matriz = [
    [1,2,3],
    [4,5,6],
    [7,8,9]
]

function imprimirTabla(){
    for (let i = 0; i < matriz.length; i++) {
        let fila = "";
        for (let j = 0; j < matriz[i].length; j++) {
            fila += matriz[i][j] + "\t";
        }
        console.log(fila);
    }
}

function sumar(){
    let suma =0
    for (let i =0; i <3 ; i ++){
        let sumafila=0
        for (let j =0; j < 3; j++){
            sumafila += matriz[i][j]
        }
        suma += sumafila
    }

    console.log("La suma total de la matriz es :", suma)
}

function intercambiar(f1,f3){
    for (let j = 0; j < 3; j++){
        let te = matriz[f1][j]
        matriz[f1][j] =  matriz[f3][j]
        matriz[f3][j] = te
    }

    imprimirTabla()

}

imprimirTabla()
sumar()
intercambiar(0,2)