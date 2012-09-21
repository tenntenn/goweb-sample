define('utils/mapper',
    [
    ],
    ()->

        # map model's properties to viewModel
        map:(viewModel, model)->
            if model
                for k, v of model
                    if viewModel[k] isnt undefined
                        vmv = viewModel[k]
                        if vmv is ko.utils.unwrapObservable(vmv)
                            viewModel[k] = ko.utils.unwrapObservable(v) 
                        else
                            viewModel[k] = ko.observable(ko.utils.unwrapObservable(v)) 

        # create model object from viewModel
        toModel:(viewModel)->
            if viewModel
                model = {}
                for k, v of viewModel
                    unwraped = ko.utils.unwrapObservable(v)
                    if viewModel.hasOwnProperty(k) and typeof unwraped isnt "function"
                        model[k] = unwraped
                        
                return model
)
