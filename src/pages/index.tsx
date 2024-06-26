import { Link } from 'umi';

const goFuncMock = (s: Function) => {
  if (s == undefined) {
    return () => {
      console.log('invoke go func mock...');
    };
  } else {
    return s;
  }
};

var query = goFuncMock(window.query);

export default function IndexPage() {
  return (
    <div>
      <h1>首页</h1>
      <Link to={'/about'}>关于</Link>
      <br />
      <button onClick={() => query('asd')}>调用go</button>
      <div id='callbackPrint'></div>
    </div>
  );
}
