var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var cos = null;
var MdEditor = null;
var headInput = null;
var ArticleItem = {};

function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}

$.ajaxSetup({
  beforeSend: setAjaxToken
});

function initEditor() {
  // 取默认标题
  headInput.val(ArticleItem.title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    editorTheme: "default",
    path: "../lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.markdown,
    codeFold: true,
    saveHTMLToTextarea: true,
    // tocm: true,
    imageUpload: true,
    taskList: true,
    // emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    imageUploadURL: "/api/v1/post-uploadfile",
    token: localStorage.getItem("AUTH_TOKEN"),
    // imageUploadCalback: function (files, cb) {
    //   uploadImage(files[0], cb);
    // },
  });
}
function uploadImage(file, cb) {
  // 创建一个新的 FormData 对象
  var formData = new FormData();

  // 将文件添加到 FormData 对象中
  formData.append('editormd-image-file', file);

  // 使用 $.ajax 发送 POST 请求
  $.ajax({
    url: "/api/v1/post-uploadfile", // 你的上传接口地址
    type: "POST",
    data: formData,
    // 必须设置正确的 Content-Type，因为 jQuery 会尝试设置它，但我们需要它保持为 multipart/form-data
    processData: false, // 告诉 jQuery 不要处理发送的数据
    contentType: false, // 告诉 jQuery 不要设置 Content-Type 请求头
    success: function(response) {
      // 假设 response 是一个包含 data.url 的 JSON 对象
      var imageUrl = response.data.url; // 直接获取 URL
      var dialog = MdEditor.imageDialog(); // 获取 image-dialog 的 DOM 元素
        dialog.find('[data-url]').val("123"); // 更新 URL 字段
      // 请求成功时调用回调函数，并传入响应数据
      cb(null, response);
    },
    error: function(jqXHR, textStatus, errorThrown) {
      // 请求失败时调用回调函数，并传入错误信息
      cb(errorThrown);
    },
    beforeSend: setAjaxToken,
  });
}


function getArticleItem(id) {
  $.ajax({
    url: "/api/v1/post/" + id,
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code != 200) {
        initEditor();
        return alert(res.error);
      }
      ArticleItem = res.data || {};
      initActive();
      initEditor();
    },
    beforeSend: setAjaxToken,
  });
}
function initActive() {
  $(".category li[value=" + ArticleItem.categoryId + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".type-box li[value=" + ArticleItem.type + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".slug-input").val(ArticleItem.slug);
}
function initCache() {
  headInput = $(".header-input");
  var query = new URLSearchParams(location.search);
  var _id = query.get("id");
  if (_id) return getArticleItem(_id);
  // 取本地缓存
  ArticleItem.title = window.localStorage.getItem(TITLE_KEY);
  ArticleItem.markdown = window.localStorage.getItem(CONTENT_KEY) || "";
  // initEditor
  initEditor();
  // 自动保存
  setInterval(() => saveHandler, AUTO_SAVE_TIME);
}

function saveHandler() {
  window.localStorage.setItem(TITLE_KEY, headInput.val());
  window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
}
function clearHandler() {
  window.localStorage.removeItem(TITLE_KEY);
  window.localStorage.removeItem(CONTENT_KEY);
}

// 发布
function publishHandler() {
  if (!ArticleItem.categoryId) return $(".publish-tip").text("请选择分类");
  ArticleItem.slug = $(".slug-input").val();
  if (ArticleItem.type == 1 && !ArticleItem.slug)
    return $(".publish-tip").text("请输入自定义链接");
  ArticleItem.title = headInput.val();
  if (!ArticleItem.title) return $(".publish-tip").text("请输入标题");
  ArticleItem.markdown = MdEditor.getMarkdown();
  if (!ArticleItem.markdown) return $(".publish-tip").text("正文");
  ArticleItem.content = MdEditor.getPreviewedHTML();

  $.ajax({
    url: "/api/v1/post",
    type: ArticleItem.pid ? "PUT" : "POST",
    contentType: "application/json",
    data: JSON.stringify(ArticleItem),
    success: function (res) {
      if (res.code !== 200) return alert(res.error);
      if (ArticleItem.pid) return $(".publish-tip").text("更新成功");
      ArticleItem = res.data || {};
      if (!ArticleItem.pid) {
        clearHandler();
      }
      location.search = "?id=" + ArticleItem.pid;
    },
    beforeSend: setAjaxToken,
  });
}

$(function () {
  // 初始化缓存
  initCache();
  // 返回首页
  var back = $(".home-btn");
  back.click(function () {
    saveHandler();
    location.href = ArticleItem.pid ? "/post/" + ArticleItem.pid : "/index.html";
  });
  if (location.search) back.text("查看");
  // 保存
  $(".save-btn").click(saveHandler);
  var drop = $(".publish-drop");
  // 显示
  $(".publish-show").click(function () {
    drop.show();
  });
  // 隐藏
  $(".publish-close").click(function () {
    drop.hide();
  });
  $(".cancel-btn").click(function () {
    drop.hide();
  });
  // 发布逻辑
  $(".publish-btn").click(publishHandler);
  // 选择分类
  $(".category").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.categoryId = target.attr("value");
    $(".publish-tip").text("");
  });
  // 选择类型
  ArticleItem.type = Number(0);
  $(".type-box").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.type = Number(target.attr("value") || 0);
  });
});
