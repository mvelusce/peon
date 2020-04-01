"""setup module"""

import setuptools

setuptools.setup(name='module_a',
                 version='1.2.3',
                 description='Module a',
                 url='some-repo',
                 author='sky',
                 zip_safe=False,
                 long_description_content_type='text/markdown',
                 packages=setuptools.find_packages(),
                 include_package_data=True,
                 install_requires=[
                     'mockito==1.1.1',
                     'paramiko==2.6',
                     'pypika==0.27.1',
                     'psycopg2-binary==2.8.2',
                     'google-auth==1.6.3',
                     'google-cloud-storage==1.17.0',
                     'google-cloud-pubsub==0.44.0'
                 ],
                 test_suite='module_a.tests',
                 scripts=['bin/run_module_a.py'])