let theme = localStorage.getItem("theme");
let themes = ["dark", "light"];

document.querySelector("html").setAttribute("data-theme", theme);
document.getElementById("toggle-theme-checkbox").checked = theme === "light";

function toggleTheme() {
  theme = themes[(themes.indexOf(theme) + 1) % themes.length];
  localStorage.setItem("theme", theme);
  document.querySelector("html").setAttribute("data-theme", theme);
}
