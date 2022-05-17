// pages/gainer_certificatelist/gainer_certificatelist.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        certificatelist: [],
        medicalName: ''
    },

    getData: function () {
        let _this = this
        let medicalName = _this.data.medicalName
        wx.request({
            url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/gainer/ViewCertificate?medicalName='+medicalName,
            header: {
                "content-type": "application/json",
                "gid": app.globalData.gid
            },
            method: "POST",
            success: function (res) {
                _this.setData({
                    certificatelist: res.data.data
                })
            }
        })
    },

    showDetailsCertificate:function(e){
        let certificate = e.currentTarget.dataset.certificate
        let medical_name = e.currentTarget.dataset.medicalname
        let address = e.currentTarget.dataset.user
        wx.navigateTo({
          url: '/pages/gainer_check/user_check?certificate='+certificate+"&medical_name="+medical_name+"&address="+address
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.setData({
            medicalName: options.MedicalName
        })
        if (this.data.medicalName != '') {
            this.getData()
        }
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    }
})