import WxValidate from '../../utils/WxValidate.js'

const date = new Date();
var year = date.getFullYear();
var month = date.getMonth() + 1;
var day = date.getDate();
if (month < 10) {undefined
month = "0" + month;
}
if (day < 10) {undefined
day = "0" + day;
}
var nowDate = year + "-" + month + "-" + day;
Page({
    /**
     * 页面的初始数据
     */
    data: {
      sexArr : ['男','女'],
      nowDate : nowDate,
      time: '获取验证码',
      count_down: 61,
      suffix: '',
      form: {
        email: '',
        code: '',
        name: '',
        password: '',
        cpassword: '',
        sex: ''
      }
    },

  bindDateChange: function (e) {
    this.setData({
     dates: e.detail.value
    })
  },
  getVerificationCode() {
    let _this = this;
    if (!_this.data.disabled) {
      _this.getCode();
    }
  },

  getCode: function(e){
    let _this = this;
    let interval = null;
    let count_down = _this.data.count_down;
    wx.showToast({
        title: '发送成功',
        icon: 'success',
        duration: 2000
      })
    interval = setInterval(function(){
      count_down--;
      _this.setData({
        time: count_down,
        suffix: "秒后可重新获取",
        disabled: true
      })
      if(count_down <= 0){
        clearInterval(interval);
        _this.setData({
          time: '重新发送',
          suffix: "",
          count_down: 61,
          disabled: false
        })
      }
    },1000)
    
    
    
  },

  deal: function(e){
    wx.navigateTo({
      url: '/pages/user_deal/user_deal',
    })
  },

  bindSexChange: function (e) {
    let sexArr = this.data.sexArr
    this.setData({
      sex : sexArr[e.detail.value]
    })
  },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
      this.initValidate();
    },
    initValidate: function(){
      const rules = {
        email:{
          required: true,
          email: true
        },
        code:{
          required: true
        },
        password:{
          required: true,
          minlength: 6,
          maxlength: 16
        },
        cpassword:{
          required: true,
          equalTo: 'password'
        },
        name:{
          required: true,
          minlength: 2,
          maxlength: 15
        },
        sex:{
          required: true
        },
        dates:{
          required: true
        },
        assistance:{
          required: true
        }
      }

      const messages = {
        email:{
          required: "请填写邮箱地址",
          email: "请填写正确的邮箱地址"
        },
        code:{
          required: "请填写验证码"
        },
        password:{
          required: "请设置密码",
          minlength: "长度不能小于6位",
          maxlength: "长度不能超过16位"
        },
        cpassword:{
          required: "请再次输入密码",
          equalTo: "两次输入密码不一致"
        },
        name:{
          required: "请填写姓名",
          minlength: "姓名长度不能小于2位",
          maxlength: "姓名长度不能超过15位"
        },
        sex:{
          required: "请选择性别",
        },
        dates:{
          required: "请选择出生日期"
        },
        assistance:{
          required: "请勾选 《A11Smile平台用户协议》",
          assistance: true
        }
      }

      this.WxValidate = new WxValidate(rules, messages)

      this.WxValidate.addMethod('assistance', (value, param) => {
        return this.WxValidate.optional(value) || (value.length >= 1 && value.length <= 2)
      }, '请勾选 《A11Smile平台用户协议》')
    },
    formSubmit: function (e) {
      console.log('form发生了submit事件，携带的数据为：', e.detail.value)
      const params = e.detail.value
       e.detail.value.osscation_address = this.data.osscation_address
       e.detail.value.date = this.data.date
       console.log(e.detail.value)
      //校验表单
      if (!this.WxValidate.checkForm(params)) {
        const error = this.WxValidate.errorList[0]
        this.showModal(error)
        return false
      }
      //向后台发送时数据 wx.request...
   
      wx.showToast({
        title: '注册成功',
        icon: 'success',
        duration: 1000,
        success: function(){
          setTimeout(
            function () {
              wx.navigateTo({
              url: '/pages/user_login/user_login',
              })
            },1000)
        }
      })
   },
   
    /***报错 **/
    showModal(error) {
       wx.showModal({
         content: error.msg
       })
    },
    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})