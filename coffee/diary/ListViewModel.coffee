define('diary/ListViewModel',
    [
        'common/settings'
        'diary/ViewModel',
    ],
    (settings, ViewModel)->
        class ListViewModel
            constructor:->
                @diaries = ko.observableArray()
                @selected = ko.observable(null)

            select: (diary)->
                if diary
                    @selected(diary)
                
            isSelected: (diary)=>
                if @selected()
                    diary is @selected()
                else
                    false

            newDiary: ()=>
                diary = new ViewModel(
                    title: '新しい日記'
                )
                @diaries.push(diary)
                diary.create()
                @select(diary)

             delDiary: ()=>
                if @selected()?
                    @selected().delete()
                    nowSelected = @diaries.indexOf(@selected())
                    @diaries.remove(@selected())
                    if nowSelected < @diaries().length
                        @select(@diaries()[nowSelected])
                    else if @diaries().length > 0
                        @select(@diaries()[nowSelected - 1])
                    else
                        @selected(null)

            loadDiaries: ()=>
                @diaries.removeAll()
                $.ajax("http://#{settings.host}/diary",
                    crossDomain:true
                    type:"GET"
                    dataType: "jsonp"
                    success:(data, dataType)=>
                        for model in data.D
                            diary = new ViewModel(model)
                            @diaries.push(diary)
                        if @diaries().length > 0
                            @select(@diaries()[0])
                )
)


