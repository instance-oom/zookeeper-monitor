$(document).ready(function () {
  $('#ServerForm')
    .form({
      on: 'blur',
      inline: true,
      fields: {
        "server_name": {
          identifier: 'Name',
          rules: [
            {
              type: 'empty',
              prompt: 'Server name cannot be null or empty'
            }
          ]
        },
        "server_ip": {
          identifier: 'IP',
          rules: [
            {
              type: 'empty',
              prompt: 'Server ip cannot be null or empty'
            },
            {
              type: 'regExp[/^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[[1-9][0-9]|[0-9])\.){3}((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[[1-9][0-9]|[1-9]))$/]',
              prompt: 'Invalidate IP.'
            }
          ]
        },
        "server_port": {
          identifier: 'Port',
          rules: [
            {
              type: 'empty',
              prompt: 'Server port cannot be null or empty'
            },
            {
              type: 'integer[1000..30000]',
              prompt: 'Server port must be integer between 1000 and 30000'
            }
          ]
        }
      }
    })

  $('#ClusterForm')
    .form({
      on: 'blur',
      inline: true,
      fields: {
        "cluster_name": {
          identifier: 'Name',
          rules: [
            {
              type: 'empty',
              prompt: 'Cluster name cannot be null or empty'
            }
          ]
        }
      }
    })

  $("#MailForm")
    .form({
      on: 'blur',
      inline: true,
      fields: {
        "address": {
          identifier: 'Address',
          rules: [
            {
              type: 'empty',
              prompt: 'Address cannot be null or empty'
            }
          ]
        }
      }
    })

  $('table.ui.sortable').tablesorter({ sortList: [[0, 0]] });
})

function delCluster(clusterID) {
  $('#confirmBox').modal({
    blurring: true,
    closable: false,
    onApprove: function () {
      var data = {
        "ID": clusterID
      }
      $.ajax({
        url: "/cluster/delete",
        data: JSON.stringify(data),
        type: "DELETE"
      }).done(function (result) {
        location.reload()
      }).fail(function (result) {
        $("#alertError").fadeIn("slow")
        $("#alertError").text(result.msg)
      }).always(function () {
        setTimeout(function () {
          $(".ui.message").fadeOut("slow")
        }, 3000)
      })
    }
  }).modal('show');
}

function delServer(serverID) {
  $('#confirmBox').modal({
    blurring: true,
    closable: false,
    onApprove: function () {
      var data = {
        "ID": serverID
      }
      $.ajax({
        url: "/server/delete",
        data: JSON.stringify(data),
        type: "DELETE"
      }).done(function (result) {
        location.reload()
      }).fail(function (result) {
        $("#alertError").fadeIn("slow")
        $("#alertError").text(result.msg)
      }).always(function () {
        setTimeout(function () {
          $(".ui.message").fadeOut("slow")
        }, 3000)
      })
    }
  }).modal('show');
}

function initBarCharts(id, title, yTitle, xData, yData) {
  $('#' + id).highcharts({
    chart: {
      type: 'column'
    },
    title: {
      text: title
    },
    xAxis: {
      categories: xData,
    },
    yAxis: {
      min: 0,
      title: {
        text: yTitle
      }
    },
    legend: {
      enabled: false
    },
    series: [{
      name: title,
      data: yData
    }],
    credits: false
  });
}

function initSplineChart(id, title, yTitle, xData, series) {
  $('#' + id).highcharts({
    chart: {
      type: 'area'
    },
    title: {
      text: title
    },
    xAxis: {
      categories: xData,
    },
    yAxis: {
      min: 0,
      title: {
        text: yTitle
      }
    },
    plotOptions: {
      spline: {
        marker: {
          enabled: true
        }
      }
    },

    series: series,
    credits: false
  });
}