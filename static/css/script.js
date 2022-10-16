async function doFunction(obj) {
    console.log(obj)
    const response = await fetch('/delete', {
        method: 'POST', // или 'PUT'
        body: JSON.stringify(obj), // данные могут быть 'строкой' или {объектом}!
        headers: {
            'Content-Type': 'application/json'
        }
    });
    window.location.reload();
}
async function doclick(obj) {
    const response = await fetch('/click', {
        method: 'POST', // или 'PUT'
        body: JSON.stringify(obj), // данные могут быть 'строкой' или {объектом}!
        headers: {
            'Content-Type': 'application/json'
        }
    });
    window.location.reload();

}
async function edit(obj){
    localStorage.clear()
    localStorage.setItem('editobj', JSON.stringify(obj));
    window.location.href = "/edit";
}
async function editDb(){
    obj = localStorage.getItem('editobj')
    objparse = JSON.parse(obj)
    localStorage.clear()
    objparse.NewCode = document.getElementById("myTextarea").value;
    console.log(objparse)
    const response = await fetch('/editdb', {
        method: 'POST', // или 'PUT'
        body: JSON.stringify(objparse), // данные могут быть 'строкой' или {объектом}!
        headers: {
            'Content-Type': 'application/json'
        }
    });
    window.location.replace("/")

}


