let large = false;
setInterval(() => {
    document.getElementById("text").style.fontSize = large ? "20px" : "40px";
    large = !large;
}, 2000);
