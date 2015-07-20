function init() {

    $.getJSON('/list', function(data) {
        json = data;
        for (item in json) {

            var i = json[item]
            console.log(i, item);
            var stp = sprintf('<li id="%s">%s<a href="#" onclick=Remove("%s");>&#x2717;</a></li>', i, item, i);
            $('ol').append(stp)
        }
    });

}


function AddItem() {
    var item = $("input").val();
    var id = makeid();
    var stp = sprintf('<li id="%s">%s<a href="#" onclick=Remove("%s");>&#x2717;</a></li>', id, safe_tags_regex(item), id);
    $('ol').append(stp)
    $.post("/task", {
        "id": id,
        "item": item
    });
}

function Remove(child) {
    $.post("/delete", {
        "id": child
    });
    $('#' + String(child)).remove();
}

/*LIBS*/
//String formatting
function sprintf() {
        var args = arguments,
            string = args[0],
            i = 1;
        return string.replace(/%((%)|s|d)/g, function(m) {
            // m is the matched format, e.g. %s, %d
            var val = null;
            if (m[2]) {
                val = m[2];
            } else {
                val = args[i];
                // A switch statement so that the formatter can be extended. Default is %s
                switch (m) {
                    case '%d':
                        val = parseFloat(val);
                        if (isNaN(val)) {
                            val = 0;
                        }
                        break;
                }
                i++;
            }
            return val;
        });
    }
    //Id gen
function makeid() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (var i = 0; i < 5; i++)
        text += possible.charAt(Math.floor(Math.random() * possible.length));
    return text;
}


function safe_tags_regex(str) {
   return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  }