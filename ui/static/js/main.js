var navLinks = document.querySelectorAll("nav a");
console.log(navLinks);
navLinks.forEach(node => {
  console.log(node)
  if (node.getAttribute('href') == window.location.pathname) {
    node.classList.add("live");
  }
});


