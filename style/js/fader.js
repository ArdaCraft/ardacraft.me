function Fader(interval, fade, urls) {
    this.interval = interval;
    this.fade = fade;
    this.top = "fader-top";
    this.bottom = "fader-bottom";
    this.urls = urls;

    this.start = function() {
        if (this.top === "" || this.bottom === "") {
            return
        }

        var btm = document.getElementById(this.bottom);
        var tp = document.getElementById(this.top);

        if (btm === undefined) {
            console.warn('Bottom element id not found!');
            return
        }
        if (tp === undefined) {
            console.warn('Top element id not found!');
            return
        }

        tp.style.opacity = 0;
        tp.style.transition = 'opacity ' + this.fade + 's ease-in-out';
        btm.style.opacity = 1;
        btm.style.transition = 'opacity ' + this.fade + 's ease-in-out';

        runloop(0, this.urls, btm, tp, this.interval, this.fade);
    };
}

function fadeText(id) {
    var text = document.getElementById(id);
}

function runloop(pos, urls, bottom, top, interval) {

    if (pos >= urls.length) {
        pos = 0;
    }

    lazyload(urls[pos], bottom, top);

    window.setTimeout(function() {
        runloop(pos + 1, urls, bottom, top, interval);
    }, interval);
}

function lazyload(url, bottom, top) {
    var downloader = new Image();
    downloader.onload = function() {
        fade(this.src, bottom, top);
    };
    downloader.src = url;
}

function fade(src, bottom, top) {
    if (top.style.opacity > 0) {
        bottom.style.backgroundImage = "url('" + src + "')";
        top.style.opacity = 0;
        bottom.style.opacity = 1;
    } else {
        top.style.backgroundImage = "url('" + src + "')";
        top.style.opacity = 1;
        bottom.style.opacity = 0;
    }
}
