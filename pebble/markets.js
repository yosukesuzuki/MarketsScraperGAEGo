var index = 0;

function showData(index,json_data){
  simply.style('large');
  simply.title(json_data.results[index].Title);
  simply.subtitle(json_data.results[index].Price);
  simply.body(json_data.results[index].Diff+'('+json_data.results[index].PriceTime+')');
}

ajax({ url: 'http://marketsapi.appspot.com/api/Markets' }, function(data){
  var json_data = JSON.parse(data);
  //console.log(json_data);
  showData(index,json_data);
 
  simply.on('singleClick', function(e) {
    if(e.button == 'up'){
      index = index - 1;
    }else if(e.button == 'down'){
       index = index + 1;
    }
    if(index >= json_data.results.length){
      index = index - json_data.results.length;
    }else if(index < 0){
      index = index +  json_data.results.length;
    }
    showData(index,json_data);
  });
  
});
