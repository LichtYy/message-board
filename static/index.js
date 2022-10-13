// 工具栏鼠标进入变色效果
var tbitems = document.getElementsByClassName("toolbaritem");
for (let i = 0; i < tbitems.length; i++) {
    tbitems[i].addEventListener('mouseenter', function () {
        this.setAttribute('style', 'color: white;background-color: gray;');
    })
    tbitems[i].addEventListener('mouseleave', function () {
        this.setAttribute('style', 'color: black;background-color: none;');
    })
}