define('diary/main',
    [
        'diary/ListViewModel'
        'diary/ViewModel'
    ],
    (ListViewModel, ViewModel)->
        $$ =
            ListViewModel: ListViewModel
            ViewModel: ViewModel
)
