pragma solidity ^0.8.0;



 import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol";  

 import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol"; 



contract A11Smile is ERC20 {

  address public A11Smile1;

  constructor() ERC20("AS","AS") { 

     A11Smile1 = msg.sender;

    _mint(msg.sender,20000 * 10 ** uint(1));

  }

//代币地址

address ercaddress = address(this);

//设置代币地址

function A11Smile_setErc(address ercaddress_) public{

     ercaddress=ercaddress_;

}



  //初始化由A11Smile给用户ETH

function A11SmileGiveETH(address payable Startuse) public payable  {

  require(msg.sender == A11Smile1, "you not can transfer");

  require(msg.sender.balance > msg.value,"Your account doesn't have so much ETH");

  Startuse.transfer(msg.value);

  emit A11GiveETH(msg.sender,Startuse,msg.value);

}





   event Addtokens( address indexed owner, uint indexed amount);//增加代币事件

   event A11GiveETH(address indexed user1, address indexed user2,uint indexed price);//A11Smile给新用户发ETH事件





//增加代币

function mint1( uint amount) public {

   _mint(msg.sender,amount ** uint(18));

   emit Addtokens(msg.sender,amount);

}



}





//征求者合约

contract UploadMedicalrecords_gainer {

   event Uploadmedicaldata(address indexed user,string indexed route,address indexed soliciter ); //用户上传医疗信息事件

   event gainerUploadmedicaldata(string indexed route,address indexed soliciter ); //征求者上传医疗信息事件

   event reward(address indexed owner,string indexed route,uint indexed price); //征求者奖励事件

//征求者地址

struct Soliciter{

  address soliciter;

  bool exist;

}

//mapping对应征求者

mapping(address => Soliciter) public Adoctor; 



//审核医疗数据时间

uint public examineTime;



//modified审核时间

modifier gainer_examineTime(address soliciter){

  examineTime = block.timestamp;



  _;

}



//设置征求者 

address[] A11soliciter;//存所有发布征求的征求者

function gainer_setDoctor(address soliciter) public{

  require(!gainer_isDoctor(soliciter),"Soliciter is exit");

  Adoctor[soliciter].soliciter=soliciter;

  Adoctor[soliciter].exist=true;

  A11soliciter.push(soliciter);

}

//查询征求者是否存在

function gainer_isDoctor(address soliciter)public view returns(bool){

  return Adoctor[soliciter].exist;

}



 string[] UpMedicalNames;

//征求者发布医疗信息

mapping(address =>string[]) UpMedicalName;//存储征求者征求的医疗数据名称

mapping(address => AddName) OnlyName;//征求者地址对应医疗数据名称

// mapping(address=>uint)A;

struct AddName {

  mapping(string =>gainer_upMedicalInformation) Addnames;//医疗数据名称对应首页展示

  mapping(string => gainer_upMedicalInformation1) Addneirong;//医疗数据名称对应内容

}

//首页展示的结构体

struct gainer_upMedicalInformation{

  // address Own;

  string MedicalName;

  uint min;

  uint max;

  uint account;

  string HospitalName;

  bool exit;

  }

//展示内容的结构体

struct gainer_upMedicalInformation1{

  uint min;  //设置最小奖励金额

  uint max;  //设置最大奖励金额

  string Medicalrecordrequirements;//病历要求

  string Requirementdescription;//需求描述

}



//征求者发布医疗信息方法

function gainer_AddMedicalInformation(string memory hospitalName_,uint min_,uint max_,uint account_,string memory MedicalName_,string memory _Medicalrecordrequirements,string memory _Requirementdescription) public {

    //   A[msg.sender]+=1;

     UpMedicalNames.push(MedicalName_);

      UpMedicalName[msg.sender].push(MedicalName_);

      OnlyName[msg.sender].Addnames[MedicalName_].MedicalName=MedicalName_;

      OnlyName[msg.sender].Addnames[MedicalName_].min=min_;

      OnlyName[msg.sender].Addnames[MedicalName_].max=max_;

      OnlyName[msg.sender].Addnames[MedicalName_].account=account_;

      OnlyName[msg.sender].Addnames[MedicalName_].HospitalName=hospitalName_;

      OnlyName[msg.sender].Addnames[MedicalName_].exit=true;

      OnlyName[msg.sender].Addneirong[MedicalName_].min=min_;

      OnlyName[msg.sender].Addneirong[MedicalName_].max=max_;

      OnlyName[msg.sender].Addneirong[MedicalName_].Medicalrecordrequirements= _Medicalrecordrequirements;

      OnlyName[msg.sender].Addneirong[MedicalName_].Requirementdescription= _Requirementdescription;

      emit gainerUploadmedicaldata(MedicalName_,msg.sender);

}





// 首页展示征求者征求的医疗数据

function SeeGainerMedicalInformationsName() public view  returns(gainer_upMedicalInformation[] memory){

 gainer_upMedicalInformation[] memory SY2 = new gainer_upMedicalInformation[](UpMedicalNames.length);

   uint m =0;

  for (uint i=0;i<A11soliciter.length;i++){

   for(uint j=0;j<UpMedicalName[A11soliciter[i]].length;j++){

      gainer_upMedicalInformation memory SY;

      SY.MedicalName = OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].MedicalName;

      SY.min = OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].min;

      SY.max = OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].max;

      SY.account =OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].account;

      SY.HospitalName = OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].HospitalName;

      SY.exit = OnlyName[A11soliciter[i]].Addnames[UpMedicalName[A11soliciter[i]][j]].exit;

      SY2[m]=SY;

      m = m + 1;

   }

  }

return SY2;

}





//用户查看征求者上传的医疗数据

function SeeGainerMedicalInformations() public view returns(gainer_upMedicalInformation1[] memory){

  gainer_upMedicalInformation1[] memory MediacalName = new gainer_upMedicalInformation1[](UpMedicalNames.length);

  uint m=0;

  for (uint i = 0; i<A11soliciter.length; i++){

    for(uint j =0;j<UpMedicalName[A11soliciter[i]].length;j++){

   gainer_upMedicalInformation1 memory MedicalContent;

   MedicalContent.min =OnlyName[A11soliciter[i]].Addneirong[UpMedicalName[A11soliciter[i]][j]].min;

   MedicalContent.max =OnlyName[A11soliciter[i]].Addneirong[UpMedicalName[A11soliciter[i]][j]].max;

   MedicalContent.Medicalrecordrequirements =OnlyName[A11soliciter[i]].Addneirong[UpMedicalName[A11soliciter[i]][j]].Medicalrecordrequirements;

   MedicalContent.Requirementdescription =OnlyName[A11soliciter[i]].Addneirong[UpMedicalName[A11soliciter[i]][j]].Requirementdescription;

   MediacalName[m] = MedicalContent;

   m = m+1;

    }

  }

  return MediacalName;

}



//征求者查看自己正在征求的医疗信息

function gainer_SeeOwnMedicalInformationing()public returns(string[] memory ,uint[] memory,uint[] memory){

    string[] memory MedicalNameIS;

    uint[] memory minIS;

    uint[] memory maxIS;

  for(uint i=0;i<UpMedicalName[msg.sender].length;i++){

    if(OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].exit=true){

        MedicalNameIS[i]=OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].MedicalName;

        minIS[i] = OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].min;

        maxIS[i] = OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].max;

    }

  }

  return(MedicalNameIS,minIS,maxIS);

}



//征求者查看自己历史的征求医疗信息

function gainer_SeeOwnMedicalInformationed()public returns(string[] memory ,uint[] memory,uint[] memory){

    string[] memory MedicalNameNo;

    uint[] memory minNo;

    uint[] memory maxNo;

    for(uint256 i=0;i<UpMedicalName[msg.sender].length;i++){

    if(OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].exit=false){

        MedicalNameNo[i]=OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].MedicalName;

        minNo[i] = OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].min;

        maxNo[i] = OnlyName[msg.sender].Addnames[UpMedicalName[msg.sender][i]].max;

    }

    }

  return(MedicalNameNo,minNo,maxNo);

}



//征求者审核

struct gainer_examine{

  address soliciter;

  address[] user;

  string[] pRoute;

  bool  whether;

}





//mapping征求者审核自己的病历信息

mapping(address => gainer_examine ) public gainerExamine;



//mapping查询医疗数据是否通过

mapping(string => gainer_examine ) public userExamineString;





// //征求者查看用户上传给自己的证书

// function  gainer_doctorsee(address soliciter) public view returns(address,string[] memory,address){

//   return (gainAmedicalInformation[soliciter].user,gainAmedicalInformation[soliciter].PictureRoute,gainAmedicalInformation[soliciter].soliciter);

// }



// //征求者审核,奖励

// function gainer_whether(address _soliciter,string memory PictureRoute,bool whether,uint ercnum_) public gainer_examineTime(_soliciter){

//         gainerExamine[_soliciter].soliciter=_soliciter;

//         gainerExamine[_soliciter].user=gainerExamineString[PictureRoute].user;

//         gainerExamine[_soliciter].pRoute=PictureRoute;

//         require(_soliciter==gainerExamineString[PictureRoute].soliciter,"The medical data does not belong to the soliciter");

//        if(whether==true){

//          gainerExamine[_soliciter].whether = true;

//         gainer_transfer_AS(gainerExamineString[PictureRoute].soliciter,gainerExamineString[PictureRoute].user,ercnum_);

//        }else{

//          revert("Failed to pass the review");

//        }

//        emit reward(_soliciter,PictureRoute,ercnum_);





//征求者奖励,奖励AS代币

function gainer_transfer_AS(address from1 ,address to1,uint quantity1) public {

  ERC20(address(this)).transferFrom(from1,to1,quantity1);

}



//查看当前账户的ETH

function getUserETH() public view returns(uint) {

  return msg.sender.balance;

}



//查看当前账户的AS

function getUserAS() public view returns(uint) {

  return ERC20(address(this)).balanceOf(msg.sender);

}

event EthgetAs(address indexed user1, address indexed user2,uint indexed price);//ETH购买AS事件

// ETH购买AS                    

function EthGetAs(address payable AddEth,address payable RedEth) public payable {

    AddEth.transfer(msg.value);

   _mint1(AddEth,RedEth,msg.value);

    emit EthgetAs(msg.sender,AddEth,msg.value);

}

function _mint1(address from,address to,uint account)  internal virtual{

    uint balance1 = ERC20(address(this)).balanceOf(from);

    require(balance1 >2 *account,"Not so many as");

    ERC20(address(this)).transferFrom(from,to,2*account);



  }



}





//用户合约

contract UploadMedicalrecords_user is UploadMedicalrecords_gainer{

//用户 

struct user_Users {

  address user;

  bool exit;

}



//上传医疗数据信息时间

uint public lastUpdateTime;



//用户上传体检报告

struct MedicalExaminationReportstrucr{

  mapping(string => string) file;

}



//mapping用户对应上传体检报告

mapping(address => MedicalExaminationReportstrucr) user_MedicalExaminationReportstrucrURL;

mapping(address => string[]) public user_MedicalExaminationReportstrucrName;



//用户上传体检报告

function user_UPMedicalExaminationReport(string memory name,string memory url) public {

  user_MedicalExaminationReportstrucrURL[msg.sender].file[name] = url;

  user_MedicalExaminationReportstrucrName[msg.sender].push(name);

}



//查看用户的体检报告量

function viewMedicalExaminationReportCount() public view returns(uint){

   return user_MedicalExaminationReportstrucrName[msg.sender].length;

}



//用户上传病历信息

struct Medicalinformation{

   mapping(string => string) file;

}

//mapping用户对应上传病历信息

mapping(address => Medicalinformation) user_MedicalinformationURL;

mapping(address => string[]) public user_MedicalinformationrName;

//用户上传病历信息

function user_UPMedicalinformation(string memory name,string memory url) public {

    user_MedicalinformationURL[msg.sender].file[name] = url;

    user_MedicalinformationrName[msg.sender].push(name);

}

//查看用户的病历信息量

function viewMedicalinformationCount() public view returns(uint){

   return user_MedicalinformationrName[msg.sender].length;

}



//用户上传证书

mapping(address=>string[])User_PictureRoute;

//用户上传证书结构体

struct user_MedicalCertificate{

  address user;

  mapping(string =>user_MedicalCertificateState)MediacalNameToSoliciter;

}

struct user_MedicalCertificateState{

  bool state;

  address soliciter;

}



//征求者查询用户上传医疗信息结构体

struct gainer_user_MedicalCertificate{

  address[] user;

  string[] PictureRoute;

  mapping(string => bool) PassOrNot;

}

//mapping用户对应上传医疗数据信息

mapping (address => user_MedicalCertificate) public userAmedicalInformation;

//征求者查询用户上传的医疗信息

mapping(address => gainer_user_MedicalCertificate)  gainerAmedicalInformation;



//modified上传时间

modifier user_upTime(address _user){

  lastUpdateTime = block.timestamp;

  _;

}



//用户上传证书

function user_AddMedicalInformation(string memory Proute,address _soliciter) public user_upTime(msg.sender) {

      require(UploadMedicalrecords_gainer.gainer_isDoctor(_soliciter),"Soliciter no exist");

<<<<<<< HEAD
      PictureRoute[msg.sender].push(Proute);

      userAmedicalInformation[msg.sender].user=msg.sender;

      userAmedicalInformation[msg.sender].Stagnation=true;

      userAmedicalInformation[msg.sender].MediacalNameToSoliciter[Proute]=_soliciter;
=======
      User_PictureRoute[msg.sender].push(Proute);

      userAmedicalInformation[msg.sender].user=msg.sender;

<<<<<<< HEAD
​      userAmedicalInformation[msg.sender].MediacalNameToSoliciter[Proute].state=true;

​      userAmedicalInformation[msg.sender].MediacalNameToSoliciter[Proute].soliciter=_soliciter;
>>>>>>> c7bc7315fa2bfc86072847e7e3dbce0a44c2e084

}



//用户查看证书是否审核true为待审核 false为审核

function user_SeeCertificateState() public view returns(string[] memory,string[] memory){

  uint m=0;

  uint n=0;

  string[] memory Auditing;

  string[] memory Reviewed;

  for(uint i=0;i<User_PictureRoute[msg.sender].length;i++){

​    if(userAmedicalInformation[msg.sender].MediacalNameToSoliciter[User_PictureRoute[msg.sender][i]].state==true){

​    Auditing[m] = User_PictureRoute[msg.sender][i];

​    m = m+1;

​    }

​    else{

​     Reviewed[n] =User_PictureRoute[msg.sender][i];

​     n = n +1;

​    }

  }

  return(Auditing,Reviewed);
=======
      userAmedicalInformation[msg.sender].Stagnation=true;

      userAmedicalInformation[msg.sender].MediacalNameToSoliciter[Proute]=_soliciter;
>>>>>>> 99fed7054e443c26b0f13707fa02c8d2c96e42ba

}



//征求者存放用户上传的医疗信息

function gainer_SeeuserUploadMedical()public {

  for(uint i = 0;i<A11user.length;i++){   

       for(uint j=0;j<User_PictureRoute[A11user[i]].length;j++){

<<<<<<< HEAD
​    if(userAmedicalInformation[A11user[i]].MediacalNameToSoliciter[User_PictureRoute[A11user[i]][j]].soliciter==msg.sender){
=======
    if(userAmedicalInformation[A11user[i]].MediacalNameToSoliciter[User_PictureRoute[A11user[i]][j]]==msg.sender){
>>>>>>> 99fed7054e443c26b0f13707fa02c8d2c96e42ba

      gainerAmedicalInformation[msg.sender].PictureRoute.push(User_PictureRoute[A11user[i]][j]);

      gainerAmedicalInformation[msg.sender].user.push(A11user[i]);

         }

       }

    }

}



//征求者审核与奖励

function gainer_whether(string memory PictureRoute,bool whether,uint ercnum_) public{

  

}



//mapping对应用户

mapping (address => user_Users)  private Ausers;



address[]  A11user;

//设置用户

function user_Adduser(address _user) public{

  require(!user_IsUser(_user),"User is exit");

  A11user.push(msg.sender);

    Ausers[_user].user= _user;

    Ausers[_user].exit= true;

}



//查询用户是否存在

function user_IsUser(address _user) public view returns(bool){

  return Ausers[_user].exit;

}





//用户查看用户上传的体检报告

function user_ViewMedicalExaminationReport(string memory name) public view returns(string memory){

  return user_MedicalExaminationReportstrucrURL[msg.sender].file[name];

}



//用户查看用户上传的医疗信息

function user_ViewMedicalInformation(string memory name) public view returns(string memory){

  return  user_MedicalinformationURL[msg.sender].file[name];

}



}