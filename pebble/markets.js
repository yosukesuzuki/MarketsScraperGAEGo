var index = 0;

function formatWithSpace(keta, mstring) {
  var src = mstring;
  var cnt = keta - mstring.length;
  if (cnt <= 0) return src;
  while (cnt-- > 0) src = src + " "; return src;
}

function localTime(timestring) {
    var ltime = new Date(timestring);
    return ltime.getMonth()+"/"+ltime.getDay()+" "+ltime.getHours()+":"+ltime.getMinutes()
}

function showData(index,json_data){
  simply.style('large');
  simply.title(json_data.results[index].Title);
  simply.subtitle(json_data.results[index].Price);
  var body_string = formatWithSpace(13,json_data.results[index].Diff)
  body_string += formatWithSpace(13,json_data.results[index].DiffPercent)
  //body_string += localTime(json_data.results[index].PriceTime)
  body_string += json_data.results[index].PriceTime
  simply.body(body_string);
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
