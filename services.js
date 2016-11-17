var express = require('express');
var app = express();

app.set('port', (process.env.PORT || 5000));


var router = express.Router();

router.get('/', function(req, res) {
    res.json({ message: 'BJJ-DB' });   
});

app.use('/api', router);

app.listen(app.get('port'), function() {
  console.log('Node app is running on port', app.get('port'));
});
