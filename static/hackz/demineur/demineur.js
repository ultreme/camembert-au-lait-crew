// BeBeLiOuS.fr

// variables
var x=0;
var y=0;
var w=31;
var h=21;
var fin=0;
var etat=new Array (w+1);
for(x=0;x<=w+1;x++)etat[x]=new Array(h+1);

var decouvert=new Array (w+1);
for(x=0;x<=w+1;x++)decouvert[x]=new Array(h+1);


//fonctions
function initialisation()
{
    id=0;
    //mine et vide
    for(y=1;y<=h;y++)
    {
        for(x=1;x<=w;x++)
        {
            if(Math.random()>0.90){etat[x][y]=99;}else{etat[x][y]=0;}
            decouvert[x][y]=0;
        }
    }
    //nombres
    for(y=1;y<=h;y++)
    {
        for(x=1;x<=w;x++)
        {
            if(etat[x][y]==0)
            {
                somme=0;
                if(etat[x+1][y]==99)somme+=1;
                if(etat[x-1][y]==99)somme+=1;
                if(etat[x][y+1]==99)somme+=1;
                if(etat[x][y-1]==99)somme+=1;
                if(etat[x-1][y-1]==99)somme+=1;
                if(etat[x+1][y+1]==99)somme+=1;
                if(etat[x-1][y+1]==99)somme+=1;
                if(etat[x+1][y-1]==99)somme+=1;
                etat[x][y]=somme;
            }
        }
    }

}
// affichage
function afficher()
{
    for(y=1;y<=h;y++)
    {
        for(x=1;x<=w;x++)
        {

            if(decouvert[x][y]==0)
            {
                document ["x"+x+"y"+y].src="/static/hackz/demineur/d_b.png";
            }
            if(decouvert[x][y]==1)
            {
                if(etat[x][y]==102)document ["x"+x+"y"+y].src="/static/hackz/demineur/d_n.pngfg";
                if(etat[x][y]==101)document ["x"+x+"y"+y].src="/static/hackz/demineur/d_o.png";
                if(etat[x][y]==100)document ["x"+x+"y"+y].src="/static/hackz/demineur/d_r.png";
                if(etat[x][y]==99)document ["x"+x+"y"+y].src="/static/hackz/demineur/d_m.png";
                if(etat[x][y]==0 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_0.png";
                if(etat[x][y]==1 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_1.png";
                if(etat[x][y]==2 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_2.png";
                if(etat[x][y]==3 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_3.png";
                if(etat[x][y]==4 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_4.png";
                if(etat[x][y]==5 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_5.png";
                if(etat[x][y]==6 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_6.png";
                if(etat[x][y]==7 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_7.png";
                if(etat[x][y]==8 )document ["x"+x+"y"+y].src="/static/hackz/demineur/d_8.png";
            }
            if(decouvert[x][y]==2)
            {
                document ["x"+x+"y"+y].src="/static/hackz/demineur/d_d.png";
            }

        }
    }
}
function fini(t)
{
    for(y=1;y<=h;y++)
    {
        for(x=1;x<=w;x++)
        {
            if(etat[x][y]!=99 && decouvert[x][y]==2){decouvert[x][y]=1;etat[x][y]=102;}
            if(etat[x][y]==99 && decouvert[x][y]==2){decouvert[x][y]=1;etat[x][y]=101;}
            if(etat[x][y]==99)decouvert[x][y]=1;
        }
    }
    if(t==0)document.getElementById("fin").innerHTML="<br><img src='/static/hackz/demineur/perdre.png' onclick='location.reload()'>";
    if(t==1)document.getElementById("fin").innerHTML="<br><img src='/static/hackz/demineur/gagner.png' onclick='location.reload()'>";
}
// logique recursive
function recursif(x,y)
{
    if(etat[x][y]==99 && decouvert[x][y]==0)
    {
        //perdu !!!!!
        fin=1;
        decouvert[x][y]=1;
        etat[x][y]=100;
        fini(0);
    }
    if(etat[x][y]>=0 && etat[x][y]<=8 && decouvert[x][y]==0)
    {
        decouvert[x][y]=1;
        if(etat[x][y]==0)
        {
            if(x<w)recursif(x+1,y);
            if(x>1)recursif(x-1,y);
            if(y<h)recursif(x,y+1);
            if(y>1)recursif(x,y-1);
            if(x<w && y<h)recursif(x+1,y+1);
            if(x>1 && y>1)recursif(x-1,y-1);
            if(y<h && x>1)recursif(x-1,y+1);
            if(y>1 && x<w)recursif(x+1,y-1);
        }
    }

}

// cliquer
function cliquer(x,y,e)
{
    if(fin==0)
    {
        if(e==0)recursif(x,y);
        if(e==2)
        {
            if(decouvert[x][y]==0)decouvert[x][y]=2;
            else {if(decouvert[x][y]==2)decouvert[x][y]=0;}
        }
        //verif
        ok=1
        for(y=1;y<=h;y++)
        {
            for(x=1;x<=w;x++)
            {
                if(decouvert[x][y]==0)ok=0;
                if(etat[x][y]==99 && decouvert[x][y]!=2)ok=0;
                if(etat[x][y]!=99 && decouvert[x][y]==2)ok=0;
            }
        }
        if(ok==1)
        {
            fin=1;
            fini(1);
        }

        afficher();
    }
}

// graphisme

document.write("<div id='plateau'>");
for(y=1;y<=h;y++)
{
    for(x=1;x<=w;x++)
    {
        document.write("<img name='x"+x+"y"+y+"' src='' onmouseup='cliquer("+x+","+y+",event.button)' oncontextmenu='return false;'>");
    }
    document.write("<br>");
}
document.write("</div><div id='fin'></div>");

// c'est parti !
initialisation();
afficher();
