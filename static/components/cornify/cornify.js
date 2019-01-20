;(function() {
  
HTMLImageElement.prototype.setStyle = function(style) {
    var image = this;
    for (var prop in style) {
        if (style.hasOwnProperty(prop)) {
            image.style[prop] = style[prop];
        }
    }
};

var cornify_add = function(top, left) {
    var windowHeight = window.innerHeight
      , windowWidth = window.innerWidth
      , body = document.getElementsByTagName('body')[0]
      , img = document.createElement('img')
      , top = top || Math.random() * windowHeight + 'px'
      , left = left || Math.random() * windowWidth + 'px'
      , unicorn_count = 7
      , rainbow_count = 4
      , imageHost = "https://raw.github.com/L1fescape/cornify/master/"

    // randomly select an image, either a unicorn or a rainbow
    var cornImage = imageHost + "images/" + 
      ((Math.round(Math.random())) ? 
        "unicorn_" + Math.floor((Math.random()*unicorn_count)+1) + ".gif" :
        "rainbow_" + Math.floor((Math.random()*rainbow_count)+1) + ".gif")
    img.setAttribute('src', cornImage);
    // used later as a query selector to find and clear cornified images
    img.setAttribute('class', 'cornify');
    img.setStyle({
        top: top,
        left: left,
        position: 'fixed',
        transition: "all .1s linear"
    });
    img.onmouseover = function() {
        var size = Math.random() + 0.5
          , angle = Math.random() * 15 + 1 + "deg"
          , transform = "rotate(" + angle + ") scale(" + size + "," + size + ")";
        this.style.transform = transform;
        this.style.WebkitTransform = transform;
    };
    img.onmouseout = function() {
        var transform = "rotate(0deg) scale(1, 1)";
        this.style.transform = transform;
        this.style.WebkitTransform = transform;
    };

    body.appendChild(img);
};

var cornify_clear = function() {
  var images = document.querySelectorAll(".cornify");
  for (var i = 0, j = images.length; i < j; i++) {
    images[i].parentNode.removeChild(images[i]);
  }
}


var cornify_pizzazz = function() {
  for (var i = 0, j = Math.random()*100 + 1; i < j; i++) {
    cornify.add();
  }
};

window.cornify = {
  add: cornify_add,
  clear: cornify_clear,
  pizzazz: cornify_pizzazz
};

})();
