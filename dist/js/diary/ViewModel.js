var __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; };

define('diary/ViewModel', ['common/settings', 'utils/mapper'], function(settings, mapper) {
  var ViewModel;
  return ViewModel = (function() {

    function ViewModel(model) {
      this["delete"] = __bind(this["delete"], this);

      this.create = __bind(this.create, this);

      this.save = __bind(this.save, this);
      this._id = "";
      this.title = ko.observable("");
      this.content = ko.observable("");
      mapper.map(this, model);
    }

    ViewModel.prototype.save = function() {
      var _this = this;
      return $.ajax("http://" + settings.host + "/diary/" + this._id, {
        crossDomain: true,
        type: "PUT",
        dataType: "json",
        data: $.toJSON(mapper.toModel(this)),
        error: function(jqXHR, textStatus, errorThrown) {
          return console.error(jqXHR);
        }
      });
    };

    ViewModel.prototype.create = function() {
      var _this = this;
      return $.ajax("http://" + settings.host + "/diary", {
        type: "POST",
        dataType: "json",
        data: $.toJSON(mapper.toModel(this)),
        success: function(data, dataType) {
          return _this._id = data._id;
        }
      });
    };

    ViewModel.prototype["delete"] = function() {
      return $.ajax("http://" + settings.host + "/diary/" + this._id, {
        type: "DELETE"
      });
    };

    return ViewModel;

  })();
});
