// pages/user_uploading/user_uploading.js

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

    },

        //JS文件   上传图片函数
upload_picture: function(name) {
    var that = this
    //让用户选择或拍摄一张照片
    wx.chooseImage({
      count: 1,	
      sizeType: ['original', 'compressed'],
      sourceType: ['album', 'camera'],
      success(res) {
      //选择完成会先返回一个临时地址保存备用
        const tempFilePaths = res.tempFilePaths
        console.log(tempFilePaths)
        //将照片上传至云端需要刚才存储的临时地址
        wx.cloud.uploadFile({
          cloudPath: 'images/test/'+new Date().getTime() + "_" +  Math.floor(Math.random()*1000)+'.jpg',
          filePath: tempFilePaths[0],
          success(res) {
              console.log(res.fileID)
              console.log('上传图片成功')
              wx.showToast({
                title: '上传成功',
                icon: 'success',
                duration: 1000,
                success: function(){
                  setTimeout(
                    function () {
                      wx.switchTab({
                          url: '/pages/user_data/user_data',
                      })
                    },1000)
                }
              })
          },
          fail(res) {
              wx.showToast({
                title: '上传失败',
                icon: 'fail',
                duration: 1000
              })
          }
        })
      }
    })
  },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

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