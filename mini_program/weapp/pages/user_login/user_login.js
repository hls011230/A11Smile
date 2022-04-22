import WxValidate from '../../utils/WxValidate.js'

let app = getApp();
Page({

    /**
     * 页面的初始数据
     */
    data: {
        email: '',
        password: '',
    },

    setEmail: function (e) {
        this.setData({
            email: e.detail.value
        })
    },

    setPasswd: function (e) {
        this.setData({
            password: e.detail.value
        })
    },

    login() {
        let email = this.data.email
        let passwd = this.data.password
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/login/",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json"
            },
            "method": "POST",
            "data": {
                "email": email,
                "passwd": passwd
            },
            success: function (res) {
                if (res.data.code == 0) {
                    app.globalData.uid = res.data.data.uid
                    wx.showToast({
                        title: '登录成功',
                        icon: 'success',
                        duration: 1000,
                        success: function () {
                            setTimeout(
                                function () {
                                    wx.switchTab({
                                        url: '/pages/user_homepage/user_homepage',
                                    })
                                }, 1000)
                        }
                    })
                } else {
                    wx.showToast({
                        title: '邮箱或密码错误',
                        icon: 'error',
                        duration: 1000,
                    })
                }
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.initValidate();
    },

    initValidate: function () {
        const rules = {
            email: {
                required: true,
                email: true
            },
            password: {
                required: true
            }
        }

        const messages = {
            email: {
                required: "请填写邮箱地址",
                email: "请填写正确的邮箱地址"
            },
            password: {
                required: "请输入密码"
            }
        }

        this.WxValidate = new WxValidate(rules, messages)
    },

    formSubmit: function (e) {
        const params = e.detail.value
        e.detail.value.osscation_address = this.data.osscation_address
        e.detail.value.date = this.data.date
        //校验表单
        if (!this.WxValidate.checkForm(params)) {
            const error = this.WxValidate.errorList[0]
            this.showModal(error)
            return false
        } else {
            this.login()
        }
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