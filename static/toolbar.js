// 加载css
var toolbarcss = document.createElement("style");
toolbarcss.innerHTML = '@import url("toolbar.css");';
var dh = document.getElementsByTagName("head")[0];
dh.append(toolbarcss);

// 加载工具栏（各个页面都需要）
var toolbar = document.createElement("div");
toolbar.setAttribute("id", "toolbar");
var db = document.getElementsByTagName("body")[0];
db.insertBefore(toolbar, db.firstChild);
var tbitems = [];
tbitems.push(document.createElement("div"));
tbitems[0].innerHTML = '<a href="/login">登录</a>';
tbitems[0].setAttribute("class", "toolbaritem first");
tbitems.push(document.createElement("div"));
tbitems[1].innerHTML = '<a href="/account">我的</a>';
tbitems[1].setAttribute("class", "toolbaritem");
tbitems.push(document.createElement("div"));
tbitems[2].innerHTML = '<a href="/about">关于</a>'
tbitems[2].setAttribute("class", "toolbaritem last");

// 工具栏鼠标进入变色效果
// var tbitems = document.getElementsByClassName("toolbaritem");
for (let i = 0; i < tbitems.length; i++) {
    toolbar.appendChild(tbitems[i])
    tbitems[i].addEventListener('mouseenter', function () {
        this.setAttribute('style', 'color: #fff;background-color: gray;');
    })
    tbitems[i].addEventListener('mouseleave', function () {
        this.setAttribute('style', 'color: rgba(77, 70, 70, 0.87);background-color: none;');
    })
}
// console.log(dh.innerHTML)
console.log(db.innerHTML)