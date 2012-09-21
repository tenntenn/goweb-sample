define('diary/ViewModel',
    [
        'common/settings'
        'utils/mapper'
    ],
    (settings, mapper)->
        class ViewModel
            constructor:(model)->
                @_id = ""
                @title = ko.observable("")
                @content = ko.observable("")

                mapper.map(@, model)

            save:=>
                $.ajax("http://#{settings.host}/diary/#{@_id}",
                    crossDomain:true
                    type:"PUT"
                    dataType:"json"
                    data:
                        $.toJSON(mapper.toModel(@))
                    error:(jqXHR, textStatus, errorThrown)=>
                       console.error(jqXHR)
                )
            create:=>
                $.ajax("http://#{settings.host}/diary",
                    type:"POST"
                    dataType:"json"
                    data:
                        $.toJSON(mapper.toModel(@))
                    success:(data, dataType)=>
                        @_id = data._id
                )
            delete:=>
                $.ajax("http://#{settings.host}/diary/#{@_id}",
                    type:"DELETE"
                )
)
