

function main(){
    rander();
}

function rander(){
    let app = document.getElementById("app");
    let h1 = document.createElement("h1");
    let str = "hello world"
    let txt = document.createTextNode(str);
    h1.appendChild(txt);
    app.appendChild(h1);
}

main()