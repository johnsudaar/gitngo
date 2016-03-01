$(function(){
  $("#custom_search_group").hide();
  $("#custom_search").change(function() {
    if($(this).prop("checked")){
      $("#custom_search_group").show();
    } else {
      $("#custom_search_group").hide();
    }
  });

  $("#search-form").submit(function(e){
    if($("#search_lang").val().length == 0) {
      $("#search_alert").removeClass("hide");
      e.preventDefault();
    }
  });
});


$(function () {
  data_pie = [];
  data_bars = [0,0,0,0,0,0,0];
  for(var i = 0; i < repositories.length ; i ++) {
    data_pie.push({
      name: repositories[i].repository.name,
      y: repositories[i].lines
    })
    d = new Date(repositories[i].repository.created_at);
    data_bars[d.getDay()]++;
  }
  $('#piecontainer').highcharts({
    chart: {
      plotBackgroundColor: null,
      plotBorderWidth: null,
      plotShadow: false,
      type: 'pie'
    },
    title: {
      text: 'Language repartition'
    },
    tooltip: {
      pointFormat: '{point.name}: <b>{point.percentage:.1f}%</b>'
    },
    plotOptions: {
      pie: {
        allowPointSelect: true,
        cursor: 'pointer',
        dataLabels: {
          enabled: true,
          format: '<b>{point.name}</b>: {point.percentage:.1f} %',
          style: {
            color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
          }
        }
      }
    },
    series: [{
      name: 'Repositories',
      colorByPoint: true,
      data: data_pie
      }]
    }
  );
  $('#barcontainer').highcharts({
        chart: {
            type: 'column'
        },
        title: {
            text: 'Project creation date by day'
        },
        xAxis: {
            categories: [
                'Sunday',
                'Monday',
                'Tuesday',
                'Wednesday',
                'Thursday',
                'Friday',
                'Saturday',
            ],
            crosshair: true
        },
        yAxis: {
            min: 0,
            title: {
                text: 'Project created'
            }
        },
        tooltip: {
            headerFormat: '<p style="font-size:15px">{point.key}</p>',
            pointFormat: '{point.y:.1f}',
            shared: true,
            useHTML: true
        },
        plotOptions: {
            column: {
                pointPadding: 0.2,
                borderWidth: 0
            }
        },
        legend: {
          enabled: false
        },
        series: [{
            name: 'Repos',
            data: data_bars

        }]
    });

});
