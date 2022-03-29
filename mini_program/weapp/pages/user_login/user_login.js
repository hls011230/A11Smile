import WxValidate from '../../utils/WxValidate.js'

Page({

    /**
     * 页面的初始数据
     */
    data: {
        email: '',
        password: ''
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.initValidate();
    },

    initValidate: function () {
        const rules = {
            email:{
                required: true,
                email: true
            },
            password:{
                required: true
            }
        }

        const messages = {
            email:{
                required: "请填写邮箱地址",
                email: "请填写正确的邮箱地址"
            },
            password:{
                required: "请输入密码"
            }
        }

        this.WxValidate = new WxValidate(rules, messages)
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
            title: '登录成功',
            icon: 'success',
            duration: 1000,
            success: function(){
              setTimeout(
                function () {
                  wx.switchTab({
                      url: '/pages/user_homepage/user_homepage',
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