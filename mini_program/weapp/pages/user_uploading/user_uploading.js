// pages/user_uploading/user_uploading.js

let app = getApp()

wx.cloud.init({
    env: 'prod-9gy59jvo10e0946b', //填上你的云开发环境id
    traceUser: true,
})
const db = wx.cloud.database()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        medical_name: '',
        address: '',
        min: '',
        max: '',
        medicalrecordrequirements: '',
        requirementdescription: '',
        resume: '',
        iconUrl:'',
        hospitalName: ''
    },

    selectCertificate: function(){
        let _this = this
        wx.navigateTo({
          url: '/pages/user_selectCertificate/user_selectCertificate?medical_name='+_this.data.medical_name+'&address='+_this.data.address,
        })
    },

    getData: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/showDetailsPage",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json"
            },
            "method": "POST",
            "data": {
                "medical_name": _this.data.medical_name,
                "address": _this.data.address
            },
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    min: data.Min,
                    max: data.Max,
                    medicalrecordrequirements: data.Medicalrecordrequirements,
                    requirementdescription: data.Requirementdescription,
                    resume: data.Resume,
                    iconUrl: data.IconUrl
                })
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        let _this = this
        _this.setData({
            medical_name: options.medical_name,
        })

        _this.setData({
            address: options.address,
        })

        _this.setData({
            hospitalName: options.hospital_name,
        })

        _this.getData()
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