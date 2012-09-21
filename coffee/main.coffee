define('main',
    [
        'diary/main'
    ],
    (diary)->
        diaryListViewModel = new diary.ListViewModel()
        diaryListViewModel.loadDiaries()
        ko.applyBindings(diaryListViewModel)
)
