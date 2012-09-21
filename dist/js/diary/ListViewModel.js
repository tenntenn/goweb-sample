var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

define('diary/ListViewModel', ['common/settings', 'diary/ViewModel'], function(settings, ViewModel) {
  var ListViewModel;
  return ListViewModel = (function() {

    function ListViewModel() {
      this.loadDiaries = __bind(this.loadDiaries, this);

      this.delDiary = __bind(this.delDiary, this);

      this.newDiary = __bind(this.newDiary, this);

      this.isSelected = __bind(this.isSelected, this);
      this.diaries = ko.observableArray();
      this.selected = ko.observable(null);
    }

    ListViewModel.prototype.select = function(diary) {
      if (diary) {
        return this.selected(diary);
      }
    };

    ListViewModel.prototype.isSelected = function(diary) {
      if (this.selected()) {
        return diary === this.selected();
      } else {
        return false;
      }
    };

    ListViewModel.prototype.newDiary = function() {
      var diary;
      diary = new ViewModel({
        title: '新しい日記'
      });
      this.diaries.push(diary);
      diary.create();
      return this.select(diary);
    };

    ListViewModel.prototype.delDiary = function() {
      var nowSelected;
      if (this.selected() != null) {
        this.selected()["delete"]();
        nowSelected = this.diaries.indexOf(this.selected());
        this.diaries.remove(this.selected());
        if (nowSelected < this.diaries().length) {
          return this.select(this.diaries()[nowSelected]);
        } else if (this.diaries().length > 0) {
          return this.select(this.diaries()[nowSelected - 1]);
        } else {
          return this.selected(null);
        }
      }
    };

    ListViewModel.prototype.loadDiaries = function() {
      var _this = this;
      this.diaries.removeAll();
      return $.ajax("http://" + settings.host + "/diary", {
        crossDomain: true,
        type: "GET",
        dataType: "jsonp",
        success: function(data, dataType) {
          var diary, model, _i, _len, _ref;
          _ref = data.D;
          for (_i = 0, _len = _ref.length; _i < _len; _i++) {
            model = _ref[_i];
            diary = new ViewModel(model);
            _this.diaries.push(diary);
          }
          if (_this.diaries().length > 0) {
            return _this.select(_this.diaries()[0]);
          }
        }
      });
    };

    return ListViewModel;

  })();
});
