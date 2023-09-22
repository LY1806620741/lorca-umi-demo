import './about.less';
import { Helmet } from 'umi';

export default () => {
  return (
    <div>
      <Helmet>
        <title>关于</title>
      </Helmet>
      <h1 className="about">这是关于页</h1>
      只展示lorca 与 react 交互，页面设计可以引入ant design
    </div>
  );
};