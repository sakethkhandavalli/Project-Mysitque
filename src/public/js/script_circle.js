$(document).ready(function() {

    $(".userId").html("1234567890");

    $.ajax({
      url: 'http://127.0.0.1:9040/wallet/getBalance/1234567890',
    })

    .done(function(data) {
      console.log(data);
      var object = JSON.parse(data);
      // $("#heading_0").css('display', 'none');
      // $("#heading_1").css('display', 'none');
      // $("#heading_2").css('display', 'none');
      // $("#heading_3").css('display', 'none');
      // setTimeout(function() {
        $(".my-progress-bar#circle_0").circularProgress({
            line_width: 8,
            color: "#00cc66",
            starting_position: 0, // 12.00 o' clock position, 25 stands for 3.00 o'clock (clock-wise)
            percent: 0, // percent starts from
            percentage: true,
            text: object["balance"],
        }).circularProgress('animate', 125, 2500);
        $(".progress-text").eq(0).css('display', 'none');
        setTimeout(function () {
          $(".progress-text").eq(0).show();
          $("#heading_0").show();
        }, 1000);
    //  },2000);

    setTimeout(function() {
      $(".my-progress-bar#circle_1").circularProgress({
          line_width: 8,
          color: "#00cc66",
          starting_position: 0, // 12.00 o' clock position, 25 stands for 3.00 o'clock (clock-wise)
          percent: 0, // percent starts from
          percentage: true,
          text: object["availableBalance"],
      }).circularProgress('animate', 125, 2500);
      $(".progress-text").eq(1).css('display', 'none');
      setTimeout(function () {
        $(".progress-text").eq(1).show();
        $("#heading_1").show();
      }, 1000);
    },0);

    setTimeout(function() {
      $(".my-progress-bar#circle_2").circularProgress({
          line_width: 8,
          color: "#00cc66",
          starting_position: 0, // 12.00 o' clock position, 25 stands for 3.00 o'clock (clock-wise)
          percent: 0, // percent starts from
          percentage: true,
          text: object["lockedBalance"],
      }).circularProgress('animate', 125, 2500);
      $(".progress-text").eq(2).css('display', 'none');
      setTimeout(function () {
        $(".progress-text").eq(2).show();
        $("#heading_2").show();
      }, 1000);
    },0);

    setTimeout(function() {
      $(".my-progress-bar#circle_3").circularProgress({
          line_width: 8,
          color: "#00cc66",
          // color: "#ED4340",
          starting_position: 0, // 12.00 o' clock position, 25 stands for 3.00 o'clock (clock-wise)
          percent: 0, // percent starts from
          percentage: true,
          text: object["underClearance"],
      }).circularProgress('animate', 125, 2500);
      $(".progress-text").eq(3).css('display', 'none');
      setTimeout(function () {
        $(".progress-text").eq(3).show();
        $("#heading_3").show();
      }, 1000);
    },0);


    //
    //   $(".my-progress-bar#circle_4").circularProgress({
    //       line_width: 8,
    //       color: "#00cc66",
    //       starting_position: 0, // 12.00 o' clock position, 25 stands for 3.00 o'clock (clock-wise)
    //       percent: 0, // percent starts from
    //       percentage: true,
    //       text: object["availableBalance"],
    //   }).circularProgress('animate', 125, 2500);
    //   $(".progress-text").eq(4).css('display', 'none');
    //   $("#heading_4").show();
    //   setTimeout(function () {
    //     $(".progress-text").eq(4).show();
    //   }, 1000);
    })

    .fail(function() {
      console.log("error");
    })
    .always(function() {
      console.log("complete");
    });

    $.ajax({
      url: 'http://127.0.0.1:9040/wallet/listAssociates/1234567890',
      type: 'GET',
    })
    .done(function(data) {
      console.log(data);
      build_table();
    })
    .fail(function() {
      console.log("error");
    })
    .always(function() {
      console.log("complete");
    });

    var build_table = function(table_info)
    {
        console.log("I entered");
        table_info = table_info.split("}");
        var parent = $(".table");
        var tbody = $("#tbody");
        for(var i=0 ; i<6; i++)
        {
          var object = table_info[i].substr(1);
          object += '}';
          var object = JSON.parse(object);
          console.log(object);
          var row = $("<tr></tr>");
          var key,j=0;
          for(key in object)
          {
            var cell = $("<td></td>");
            if(j==0)
            {
              var anch = $("<a></a>").attr("href", "/trackingPage");
              anch.text(object[key]);
              cell.append(anch);
            }
            else {
              cell.text(object[key]);
            }
            row.append(cell);
            j+=1;
          }
          tbody.append(row);
        }
        parent.append(tbody);
    };

});


// navbar

$(function(){
    // $('[data-toggle="tooltip"]').tooltip();
    $(".side-nav .collapse").on("hide.bs.collapse", function() {
        $(this).prev().find(".fa").eq(1).removeClass("fa-angle-right").addClass("fa-angle-down");
    });
    $('.side-nav .collapse').on("show.bs.collapse", function() {
        $(this).prev().find(".fa").eq(1).removeClass("fa-angle-down").addClass("fa-angle-right");
    });
})

$("#menu-toggle").click(function(e) {
        //e.preventDefault();
        console.log("working");
        $("#collapse_wrapper").toggleClass("active");
});

var addAssociatesJs = function()
{
  console
  var associateID = $("#associateID").val();
  var associateName = $("#associateName").val();
  var associateWalletAddress = $("#associateWalletAddress").val();
  $.ajax({
    url: 'http://127.0.0.1:9040/wallet/addAssociates/1234567890',
    type: 'POST',
    dataType: 'json',
    data: JSON.stringify({"associateID" : String(associateID),
            "associateName" : String(associateName),
            "associateWalletAddress" : String(associateWalletAddress)}),
  })
  .done(function() {
    console.log("success");
  })
  .fail(function() {
    console.log("error");
  })
  .always(function() {
    console.log("complete");
  });

}


var addTransactionJs = function()
{
  function makeid() {
  var text = "";
  var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`,./;[]=!@#$%^&*()_+<>?:{}|";

  for (var i = 0; i < 25; i++)
    text += possible.charAt(Math.floor(Math.random() * possible.length));

  return text;
}

  var transactionId = makeid();
  var time = new Date();
  var Amount = $("#Amount").val();
  if(isNaN(Amount) === true)
  {
    alert("number raa othaalsindhi pichi pulka");
    //break;
    return;
  }
  console.log(Amount);
  var NatureOfTransaction = $("#NatureOfTransaction").val();
  function getSourceAdress(data)
  {
    data = data.split('"WalletInfo":');
    data = data[1].split(',"status"');
    return String(data);
  }
  var a = function getSourceInfo()
  {

    $.get("http://127.0.0.1:9040/wallet/1234567890", function(data, status)
    {
      getSourceAdress(data);
    }
  )};
  var SourceAddress = a();
  $.ajax({
    url: 'http://127.0.0.1:9040/wallet/addTransaction/1234567890',
    type: 'POST',
    dataType: 'json',
    data: JSON.stringify({"transactionId" : transactionId,
            "time" : time,
            "amount" : parseInt(Amount),
            "nature" : String(NatureOfTransaction),
            "sourceAddress" : SourceAddress,
            "destinationAddress" : SourceAddress}),

  })
  .done(function() {
    console.log("success");
  })
  .fail(function() {
    console.log("error");
    alert("paisal lo alphabet undavu raa ooroodaa");
  })
  .always(function() {
    console.log("complete");
  });

}
