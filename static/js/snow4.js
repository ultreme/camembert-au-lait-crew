/*
The uncompressed source looked something like this before I started shrinking it.
This isn't exactly the same though, I had to change a bit of functionality to make it fit into 1024 bytes.
I added a few comments too.
 */

// initialize the grids and canvas and pen:
var b = document.body;
var c = document.getElementById("neige");
var a = c.getContext('2d');
document.body.clientWidth; // fix bug in webkit: http://qfox.nl/weblog/218

c.height=300;
g=[], h=[];
for (x=-4; x<80; x++) {
    g[x]=[], h[x]=[];
    for (y=-1; y<76; y++) g[x][y]=h[x][y]=1;
}
p = 3;

onkeydown = function (e) {
    k = e.keyCode;
    if (48<k&&k<55) p = k-48;
};

// I later discovered that |0 is better than ~~ because it has lower precedence.
b.onmousemove = function (e) { g[Math.round(e.pageX*75/b.clientWidth)][Math.round(e.pageY*70/b.clientHeight)] = p; };

setInterval(function () {
    for (x=0; x<75; x++) {
        r = ~~(Math.random()*75);
        R = ~~(r%3-1);
        for (y=0; y<76; y++) {
            j=g[x]; J=h[x];
            t = j[y];
            d = j[y+1];
            u = j[y-1];
            s=0;

            // all the cell logic happens down here:
            if (t==5) {
                if (R&&u^4) h[x][y-1] = 5; // for integers, ^ is the same as !=
            }
            if (t==4&&h[x][y]==1) {
                if (d==1) h[x][y+1]=4;
                else if (g[x+R][y+1]==1&&g[x+R][y]^4) h[x+R][y+1]=4;
                else if (g[x-R][y+1]==1&&g[x-R][y]^4) h[x-R][y+1]=4;
                else if (g[x+R][y]==1&&h[x+R][y]==1) h[x+R][y] = 4;
                else s=1;
            }
            if (t==3) {
                // in the compressed source I used ternary operator in place of ifs as much as possible
                d==1||d==4 ? (h[x][y+1]=3, h[x][y]=d)
                    : (g[x+R][y+1]==1&&g[x+R][y]^3 ? h[x+R][y+1]=3
                       : (g[x-R][y+1]==1&&g[x-R][y]^3 ? h[x-R][y+1]=3
                          : h[x][y] = g[x+R][y+R]==5 ? 4 : t));
            }
            if (t==2) s=1;
            if (t==6) {
                if (d==3&&g[x][y+2]==3) {
                    if (u^0&&u^5&&!r%9) for (i=-4;i<5;i++) h[x+i][y-1] = g[x+i][y-1]==1 ? 0:g[x+i][y-1];
                    s=1;
                }
                else h[x][y+1] = 6;
            }

            // 0 is the cell code for tree leaves.
            if (t==0)
                g[x-1][y]==1&&g[x+1][y]==0&&g[x-2][y]^1&&R&&R ? h[x+1][y-1]=0 :0,
            g[x+1][y]==1&&g[x-1][y]==0&&g[x+2][y]^1&&R ? h[x-1][y-1]=0 :0,
            g[x-1][y]==0&&g[x+1][y]==0&&R ? h[x][y-1]=0 :0,
            s=1;

            if (s) h[x][y] = t==0||t==6? (g[x+1][y+R]==5||g[x-1][y-R]==5?5:t) : t;

            // picking colours from an array literal:
            a.fillStyle = '#'+['2d3','000','877','eff','24f','f90','860'][t];
            a.fillRect(x*4, y*4, 4, 4);
        }
    }
    for (x=0; x<75; x++) for (y=0; y<76; y++) g[x][y] = 1; // clear before swapping...
    G=g;H=h;g=H;h=G; // swap grids
    g[r][0] = 3; // add random snow to the top of the screen.
}, 75);
