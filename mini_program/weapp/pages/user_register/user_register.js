import WxValidate from '../../utils/WxValidate.js'

Page({
    /**
     * 页面的初始数据
     */
    data: {
        time: '获取验证码',
        count_down: 61,
        suffix: '',
        email: '',
        code: '',
        username: '',
        resume: '',
        password: '',
        cpassword: '',
    },

    setEmail: function (e) {
        this.setData({
            email: e.detail.value
        })
    },

    setCode: function (e) {
        this.setData({
            code: e.detail.value
        })
    },

    setPasswd: function (e) {
        this.setData({
            password: e.detail.value
        })
    },

    setUserName: function (e) {
        this.setData({
            username: e.detail.value
        })
    },

    setResume: function (e) {
        this.setData({
            resume: e.detail.value
        })
    },

    getVerificationCode() {
        let _this = this;
        let reg = /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/
        let email = _this.data.email
        if (reg.test(email)) {
            if (!_this.data.disabled) {
                _this.getCode();
            }
        } else {
            wx.showToast({
                title: '邮箱格式不正确',
                icon: 'error',
                duration: 2000
            })
        }
    },

    getCode: function (e) {
        let _this = this;
        let interval = null;
        let count_down = _this.data.count_down;
        wx.showToast({
            title: '发送成功',
            icon: 'success',
            duration: 2000
        })
        interval = setInterval(function () {
            count_down--;
            _this.setData({
                time: count_down,
                suffix: "秒后可重新获取",
                disabled: true
            })
            if (count_down <= 0) {
                clearInterval(interval);
                _this.setData({
                    time: '重新发送',
                    suffix: "",
                    count_down: 61,
                    disabled: false
                })
            }
        }, 1000)
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/register/sendEmail",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json"
            },
            "method": "POST",
            "data": {
                "email": this.data.email
            }
        })
    },

    verifyEmail() {
        let _this = this;
        let code = _this.data.code
        let email = _this.data.email
        let uname = _this.data.username
        let resume = _this.data.resume
        let passwd = _this.data.password
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/register/verifyEmail",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json"
            },
            "method": "POST",
            "data": {
                "email": email,
                "code": code
            },
            success: function (res) {
                if (!res.data.code == 0) {
                    wx.showToast({
                        title: '验证码错误',
                        icon: 'error',
                        duration: 2000
                    })
                } else {
                    wx.cloud.callContainer({
                        "config": {
                            "env": "prod-9gy59jvo10e0946b"
                        },
                        "path": "/user/register/",
                        "header": {
                            "X-WX-SERVICE": "test-allsmile",
                            "content-type": "application/json"
                        },
                        "method": "POST",
                        "data": {
                            "email": email,
                            "passwd": passwd,
                            "uname": uname,
                            "resume": resume
                        },
                        success: function (res) {
                            if (res.data.code == 0) {
                                wx.showToast({
                                    title: '注册成功',
                                    icon: 'success',
                                    duration: 1000,
                                    success: function () {
                                        setTimeout(
                                            function () {
                                                wx.navigateTo({
                                                    url: '/pages/user_login/user_login',
                                                })
                                            }, 1000)
                                    }
                                })
                            } else {
                                wx.showToast({
                                    title: '注册失败',
                                    icon: 'error',
                                    duration: 1000,
                                })
                            }
                        }
                    })
                }
            }
        })
    },

    deal: function (e) {
        wx.navigateTo({
            url: '/pages/user_deal/user_deal',
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
            code: {
                required: true
            },
            username: {
                required: true,
                maxlength: 16
            },
            password: {
                required: true,
                minlength: 6,
                maxlength: 16
            },
            cpassword: {
                required: true,
                equalTo: 'password'
            },
            assistance: {
                required: true
            }
        }

        const messages = {
            email: {
                required: "请填写邮箱地址",
                email: "请填写正确的邮箱地址"
            },
            code: {
                required: "请填写验证码"
            },
            username: {
                required: "请填写用户名",
                maxlength: "长度不能超过16位"
            },
            password: {
                required: "请设置密码",
                minlength: "长度不能小于6位",
                maxlength: "长度不能超过16位"
            },
            cpassword: {
                required: "请再次输入密码",
                equalTo: "两次输入密码不一致"
            },
            assistance: {
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
        const params = e.detail.value
        //校验表单
        if (!this.WxValidate.checkForm(params)) {
            const error = this.WxValidate.errorList[0]
            this.showModal(error)
            return false
        } else {
            this.verifyEmail()
        }
        //向后台发送时数据 wx.request...
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