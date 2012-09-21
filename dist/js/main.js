
define('main', ['diary/main'], function(diary) {
  var diaryListViewModel;
  diaryListViewModel = new diary.ListViewModel();
  diaryListViewModel.loadDiaries();
  return ko.applyBindings(diaryListViewModel);
});
