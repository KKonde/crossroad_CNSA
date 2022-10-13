const addBtn = document.querySelector("#addBtn");
const deleteBtn = document.querySelector(".deleteBtn");
const token = document.querySelector("#token");

function setButtons() {

    addBtn.addEventListener('click', () => {
        document.querySelector(".img img").remove();
        addBtn.value = "";
        token.value = "1";
    });
    addBtn.addEventListener('change', () => {
        var reader = new FileReader();
        reader.onload = function (event) {
            var img = document.createElement("img");
            img.setAttribute("src", event.target.result);
            img.style.width = "100%";
            img.style.width = "100%";
            document.querySelector(".img").appendChild(img);
        };
        reader.readAsDataURL(event.target.files[0]);
    });


    deleteBtn.addEventListener('click', () => {
        document.querySelector(".img img").remove();
        addBtn.value = "";
        token.value = "1";
    });

}


setButtons();
