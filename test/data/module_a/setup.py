"""setup module"""

import setuptools

setuptools.setup(name='skytv-worker-sftp',
                 version='1.1.32',
                 description='Worker component for SFTP sources',
                 url='https://github.com/sky-uk/ita-data-analytics',
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
                     'google-cloud-pubsub==0.44.0',
                     'skytv-base',
                     'skytv-cloud-storage',
                     'skytv-masking',
                     'skytv-pubsub',
                     'skytv-sftp',
                     'skytv-files',
                     'skytv-metamodelclient',
                     'skytv-statefulness'
                 ],
                 test_suite='worker_sftp.tests',
                 scripts=['bin/run_worker_sftp.py'])
