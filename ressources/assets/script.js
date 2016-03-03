$(function(){

  // Hide custom search form (unless the checkbox is already checked (Using back))
  if(! $("#custom_search").prop("checked")){
    $("#custom_search_group").hide();
  }

  // Hide/Show custom search form
  $("#custom_search").change(function() {
    if($(this).prop("checked")){
      $("#custom_search_group").show();
    } else {
      $("#custom_search_group").hide();
    }
  });

  // On Search form validation
  $("#search-form").submit(function(e){
    // If there was no language specified
    if($("#search_lang").val().length == 0) {
      // Show error message and prevent from continuing
      $("#search_alert").removeClass("hide");
      e.preventDefault();
    } else {
      // else, show the loading icon
      $("#main_form").addClass("hide");
      $("#loader").removeClass("hide");
    }
  });
});

// Boostrap intitialisations
$(function () {
  $('[data-toggle="tooltip"]').tooltip()
})


$(function () {
  // Data used to generate the pie chart
  data_pie = [];
  // Data used to generate the bar chart (7 days, 0 repository per day)
  data_bars = [0,0,0,0,0,0,0];

  // Filling in the data_bars array by looping over the repository tab.
  for(var i = 0; i < repositories.length ; i ++) {
    data_pie.push({
      name: repositories[i].repository.name,
      y: repositories[i].bytes
    })
    d = new Date(repositories[i].repository.created_at);
    data_bars[d.getDay()]++;
  }

  // --------------------------------------------
  // -            PIE CHART GENERATION          -
  // --------------------------------------------
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

  // --------------------------------------------
  // -            BAR CHART GENERATION          -
  // --------------------------------------------

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
