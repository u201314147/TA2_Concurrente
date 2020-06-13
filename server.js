const port = 4000;
const axios = require('axios')
var express = require('express'),
    app = express();
var bodyParser = require('body-parser');

app.use(express.static(__dirname));

app.use(bodyParser.urlencoded({
   extended: false
}));

app.use(bodyParser.json());

app.get('/', function(req, res){
  res.render('form');// if jade
  // You should use one of line depending on type of frontend you are with
 // res.sendFile(__dirname + '/form.html'); //if html file is root directory
 res.sendFile("index.html"); //if html file is within public directory
});


//const data = JSON.stringify({
//  todo: 'Buy the milk'
//})


app.post('/',function(req,res){
  var ID = Math.random().toString(36).substr(2, 9);
   var Fixedacidity = req.body.Fixedacidity;
   var Volatileacidity = req.body.Volatileacidity;
   var Citricacid = req.body.Citricacid;
   var Residualsugar = req.body.Residualsugar;
   var Chlorides = req.body.Chlorides;
   var Freesulfurdioxide = req.body.Freesulfurdioxide;
   var Totalsulfurdioxide = req.body.Totalsulfurdioxide;
   var Density = req.body.Density;
   var PH = req.body.PH;
   var Sulphates = req.body.Sulphates;
   var Alcohol = req.body.Alcohol;
   var Quality  = req.body.Quality ;



   var htmlData = "{" +
   			       '"Fixedacidity": "' + Fixedacidity +'",\n'
   				  + '"Volatileacidity": "' + Volatileacidity +'",\n'
   				  + '"Citricacid": "' + Citricacid +'",\n'
   				  + '"Residualsugar": "' + Residualsugar +'",\n'
   				  + '"Chlorides": "' + Chlorides +'",\n'
   				  + '"Freesulfurdioxide": "' + Freesulfurdioxide +'",\n'
   				  + '"Totalsulfurdioxide": "' + Totalsulfurdioxide +'",\n'
   				  + '"Density": "' + Density +'",\n'
   				  + '"PH": "' + PH +'",\n'
   				  + '"Sulphates": "' + Sulphates +'",\n'
   				  + '"Alcohol": "' + Alcohol +'",\n'
   				  + '"Quality": "' + Quality  +'",\n'
   				  +"}";


   res.send(htmlData);

   console.log(htmlData);

axios.post('http://localhost:3000/gophers', {
  ID,
    Fixedacidity,
    Volatileacidity,
   Citricacid,
   Residualsugar,
   Chlorides,
   Freesulfurdioxide,
   Totalsulfurdioxide,
   Density,
   PH,
   Sulphates,
   Alcohol,
   Quality
})
.then((res) => {
  console.log(`statusCode: ${res.statusCode}`)
  console.log(res)
})
.catch((error) => {
  console.error(error)
})


//   res.redirect('http://localhost:3000');
});

app.listen(port,() => {
  console.log("Started on PORT 4000");
})